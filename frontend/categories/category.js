(async function initCategories() {
    "use strict";

    const API_BASE = "http://localhost:8081/categorias";
    let cats = [];
    window.cats = cats;

    function escapeHtml(str = "") {
        return String(str)
        .replace(/&/g, "&amp;")
        .replace(/</g, "&lt;")
        .replace(/>/g, "&gt;")
        .replace(/"/g, "&quot;")
        .replace(/'/g, "&#039;");
    }

    try {
        // 1) cargar template HTML
        const tplRes = await fetch("categories/category.html");
        if (!tplRes.ok) throw new Error(`Can't show: ${tplRes.status}`);
        const tplHtml = await tplRes.text();
        const container = document.getElementById("categories");
        if (!container) throw new Error("No se encontró el contenedor con id='categories' en el DOM.");
        container.innerHTML = tplHtml;

        // 2) obtener referencias DOM desde la plantilla inyectada
        const categoryList = document.getElementById("category-list");
        const btnAddCategory = document.getElementById("add-category");
        const modalAdd = document.getElementById("modal-category-overlay");
        const btnCloseModal = document.getElementById("modal-category-close");
        const saveCategoryBtn = document.getElementById("save-category");
        const nameInput = document.getElementById("category-name");
        const descInput = document.getElementById("category-description");

        const modalDelete = document.getElementById("modal-delete-overlay");
        const btnCloseDelete = document.getElementById("modal-delete-close");
        const btnConfirmDelete = document.getElementById("confirm-delete");
        const btnCancelDelete = document.getElementById("cancel-delete");

        if (!categoryList) throw new Error("No se encontró el elemento #category-list en la plantilla.");

        // Estado temporal para borrado
        let categoryToDeleteId = null;

        // Render de lista de categorías (recibe array)
        function renderCategories(cats = []) {
            categoryList.innerHTML = "";
            if (!cats.length) {
                categoryList.innerHTML = "<p>There are not categories.</p>";
                return;
            }

            cats.forEach(cat => {
                const div = document.createElement("div");
                div.className = "category";
                div.dataset.id = cat.id;

                const name = escapeHtml(cat.name ?? "");
                const description = escapeHtml(cat.description ?? "");

                div.innerHTML = `
                    <div class="category-left" style="display:flex;align-items:center;gap:10px">
                        <input type="checkbox" id="cat-${cat.id}">
                        <div style="text-align:left;">
                        <label for="cat-${cat.id}">${name}</label>
                        ${description ? `<div style="font-size:14px;color:#333;margin-top:4px">${description}</div>` : ""}
                        </div>
                    </div>
                    <span class="delete-category" title="Eliminar categoría">&times;</span>
                `;

                categoryList.appendChild(div);
            });
        }

        // Obtener categorías desde API (GET)
        async function fetchCategories() {
            categoryList.innerHTML = "<p>Loading Categories...</p>";
            try {
                const res = await fetch(API_BASE, { method: "GET" });
                if (!res.ok) throw new Error(`GET failed: ${res.status}`);
                const json = await res.json();

                if (Array.isArray(json)) cats = json;
                else if (Array.isArray(json.categorias)) cats = json.categorias;
                else if (Array.isArray(json.categories)) cats = json.categories;
                else if (json.category) {
                    cats = Array.isArray(json.category) ? json.category : [json.category];
                } else {
                    cats = Array.isArray(json) ? json : [];
                }

                cats.sort((a, b) => {
                    const ida = Number(a?.id ?? 0);
                    const idb = Number(b?.id ?? 0);
                    return idb - ida;
                });

                window.cats = cats;
                renderCategories(cats);
                window.dispatchEvent(new Event('categories:loaded'));

            } catch (err) {
                console.error(err);
                categoryList.innerHTML = "<p>Error loading categories.</p>";
            }
        }

        // Abrir modal añadir
        btnAddCategory.onclick = () => {
            if (nameInput) nameInput.value = "";
            if (descInput) descInput.value = "";
            modalAdd.style.display = "block";
            setTimeout(() => nameInput?.focus?.(), 50);
        };

        // Cerrar modal añadir
        btnCloseModal.onclick = () => {
            modalAdd.style.display = "none";
        };

        // Cerrar modales al hacer click fuera
        window.addEventListener("click", (e) => {
            if (e.target === modalAdd) modalAdd.style.display = "none";
            if (e.target === modalDelete) modalDelete.style.display = "none";
        });

        // Delegación para botón de eliminar (la X)
        categoryList.addEventListener("click", (e) => {
            if (e.target.classList.contains("delete-category")) {
                const card = e.target.closest(".category");
                categoryToDeleteId = card?.dataset?.id ?? null;
                modalDelete.style.display = "block";
            }
        });

        // Cerrar modal borrar
        btnCloseDelete.onclick = () => (modalDelete.style.display = "none");
        btnCancelDelete.onclick = () => (modalDelete.style.display = "none");

        // Confirmar borrado -> DELETE
        btnConfirmDelete.onclick = async () => {
            if (!categoryToDeleteId) {
                modalDelete.style.display = "none";
                return;
            }
            btnConfirmDelete.disabled = true;
            try {
                const res = await fetch(`${API_BASE}/${categoryToDeleteId}`, { method: "DELETE" });
                if (!res.ok) {
                const txt = await res.text().catch(() => "");
                throw new Error(txt || `DELETE failed: ${res.status}`);
                }
                await fetchCategories();
            } catch (err) {
                console.error(err);
                alert("The category could not be deleted: " + err.message);
            } finally {
                btnConfirmDelete.disabled = false;
                categoryToDeleteId = null;
                modalDelete.style.display = "none";
            }
        };

        // Guardar nueva categoría -> POST
        saveCategoryBtn.onclick = async () => {
            const name = nameInput?.value?.trim?.() ?? "";
            const description = descInput?.value?.trim?.() ?? "";

            if (!name) {
                alert("The name of the category is mandatory.");
                return;
            } else if (!description) {
                alert("The description of the category is mandatory.");
                return;
            }

            saveCategoryBtn.disabled = true;
            try {
                const res = await fetch(API_BASE, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ name, description })
                });
                if (!res.ok) {
                const txt = await res.text().catch(() => "");
                throw new Error(txt || `POST failed: ${res.status}`);
                }
                
                await fetchCategories();
                modalAdd.style.display = "none";
            } catch (err) {
                console.error(err);
                alert("The category cound not be created: " + err.message);
            } finally {
                saveCategoryBtn.disabled = false;
            }
        };

        await fetchCategories();

    } catch (err) {
        console.error("Error initialize the categories:", err);
        const c = document.getElementById("categories");
        if (c) c.innerHTML = `<p>Error initialize the categories: ${escapeHtml(err.message)}</p>`;
    }
})();