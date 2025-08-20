fetch("tasks/task.html")
    .then(res => res.text())
    .then(html => {
        document.getElementById("tasks").innerHTML = html;

        // elementos del DOM
        const modalTask = document.getElementById("modal-task-overlay");
        const modalTaskClose = document.getElementById("modal-task-close");
        const btnSaveTask = document.getElementById("save-task");
        const txtTask = document.getElementById("task-text");
        const inputDue = document.getElementById("task-due");
        const selectCategory = document.getElementById("task-category");

        const editModal = document.getElementById("modal-edit-task-overlay");
        const editClose = document.getElementById("modal-edit-task-close");
        const editText = document.getElementById("edit-task-text");
        const editDue = document.getElementById("edit-task-due");
        const editStatus = document.getElementById("edit-task-status");
        const btnDelete = document.getElementById("edit-task-delete");
        const btnSave = document.getElementById("edit-task-save");

        let currentEditTaskId = null;
        let originalEditSnapshot = null;

        // Estado dinámico: ahora sí es mutable
        let TASKS = [];

        // Cambia la URL base si tu API está en otro puerto/dominio
        const API_BASE = "http://localhost:8082/tareas";

        let activeCategoryFilter = new Set(); // ids numéricos o strings
        let activeStatusFilter = new Set();   // tokens: "no_started" | "started" | "ended"


        // util: leer categorías seleccionadas desde el DOM (#category-list)
        function readSelectedCategoryIds() {
            const inputs = document.querySelectorAll('#category-list input[type="checkbox"][id^="cat-"]');
            const ids = [];
            inputs.forEach(i => {
                if (i.checked) {
                    // id tiene formato cat-<id>
                    const id = i.id.replace(/^cat-/, '');
                    if (id !== "") ids.push(isNaN(Number(id)) ? id : Number(id));
                }
            });
            return ids;
        }

        // util: leer estados seleccionados desde el DOM (#legend)
        function readSelectedStatuses() {
            // asumimos que los checkboxes están dentro de #legend con los contenedores
            const statuses = [];
            const noBox = document.querySelector('#legend-item-legend-no input[type="checkbox"]');
            const startedBox = document.querySelector('#legend-item-legend-started input[type="checkbox"]');
            const endedBox = document.querySelector('#legend-item-legend-ended input[type="checkbox"]');

            if (noBox && noBox.checked) statuses.push('no_started');
            if (startedBox && startedBox.checked) statuses.push('started');
            if (endedBox && endedBox.checked) statuses.push('ended');

            return statuses;
        }

        // aplica los filtros actuales sobre TASKS y renderiza
        function applyFilters() {
            // si no hay tareas cargadas, delegar render vacío
            if (!Array.isArray(TASKS)) {
                renderTasks([]);
                return;
            }

            // actualizar sets desde DOM para mantener sincronía (evita estados obsoletos)
            activeCategoryFilter = new Set(readSelectedCategoryIds());
            activeStatusFilter = new Set(readSelectedStatuses());

            // si ambos sets están vacíos => mostramos todas las tareas
            const noCategoryFilter = activeCategoryFilter.size === 0;
            const noStatusFilter = activeStatusFilter.size === 0;

            const filtered = TASKS.filter(t => {
                // categoría: si no hay filtro de categoría, pasar. Si hay, comprobar igualdad
                if (!noCategoryFilter) {
                    const catId = (t.category_id == null) ? "" : (isNaN(Number(t.category_id)) ? t.category_id : Number(t.category_id));
                    if (!activeCategoryFilter.has(catId)) return false;
                }
                // estado: si no hay filtro de estado, pasar. Si hay, comprobar token
                if (!noStatusFilter) {
                    if (!t.status) return false;
                    if (!activeStatusFilter.has(t.status)) return false;
                }
                return true;
            });

            renderTasks(filtered);
        }

        // attach: escucha cambios en los checkboxes de categorías y leyenda (delegación)
        function attachFilterListeners() {
            // delegación: escucha cambios en todo el documento para inputs de categoría
            document.addEventListener('change', (e) => {
                const target = e.target;
                if (!(target instanceof Element)) return;

                // categoría (id empieza por cat- dentro de #category-list)
                if (target.matches('#category-list input[type="checkbox"][id^="cat-"]')) {
                    // Cuando las categorías cambian, actualizar y aplicar filtros
                    activeCategoryFilter = new Set(readSelectedCategoryIds());
                    applyFilters();
                    return;
                }

                // legend (status)
                if (target.closest('#legend')) {
                    if (target.matches('#legend input[type="checkbox"], #legend input[type="checkbox"]')) {
                        activeStatusFilter = new Set(readSelectedStatuses());
                        applyFilters();
                        return;
                    }
                }
            }, { capture: false });
        }

        // inicializamos listeners de filtro
        attachFilterListeners();

        // Asegurarnos de que al cargar categorías (cuando categories module re-render) los filtros se reapliquen
        window.addEventListener('categories:loaded', () => {
            // rellenamos el select y re-aplicamos filtros en caso de que haya checkboxes ya marcados
            populateCategorySelect?.();
            applyFilters();
        });

        // normaliza acentos (compatible con la mayoría de navegadores)
        function removeDiacritics(str = "") {
            return String(str).normalize('NFD').replace(/[\u0300-\u036f]/g, '');
        }

        // convierte estado backend (español/inglés, con o sin acentos) -> token frontend
        function mapFromBackendStatus(status) {
            if (!status) return "no_started";
            const s = removeDiacritics(String(status)).toLowerCase();

            // coincidencias tolerantes (maneja "Pendiente", "En Progreso", "Finalizado"
            // y variantes en inglés si acaso: "Pending", "In Progress", "Finished")
            if (s.includes("pend") || s.includes("pending")) return "no_started";
            if (s.includes("en progreso") || s.includes("enprogreso") || s.includes("progres") || s.includes("in progress")) return "started";
            if (s.includes("finaliz") || s.includes("finalizad") || s.includes("finished") || s.includes("final")) return "ended";

            // fallback: si no encaja, devolver 'no_started'
            return "no_started";
        }

        // convierte token frontend -> texto que espera el backend
        function mapToBackendStatus(frontStatus) {
            switch (frontStatus) {
                case "no_started": return "Pendiente";
                case "started":    return "En Progreso";
                case "ended":      return "Finalizado";
                default:           return "Pendiente";
            }
        }

        function formatDate(iso) {
            if (!iso) return "";
            const d = new Date(iso);
            const dd = String(d.getDate()).padStart(2,'0');
            const mm = String(d.getMonth()+1).padStart(2,'0');
            const yyyy = d.getFullYear();
            return `${dd}/${mm}/${yyyy}`;
        }

        function isoToLocalDatetime(iso) {
            if (!iso) return "";
            const d = new Date(iso);
            const y = d.getFullYear();
            const m = String(d.getMonth() + 1).padStart(2, '0');
            const day = String(d.getDate()).padStart(2, '0');
            const hh = String(d.getHours()).padStart(2, '0');
            const mm = String(d.getMinutes()).padStart(2, '0');
            return `${y}-${m}-${day}T${hh}:${mm}`;
        }

        function hasAuthToken() {
            return !!localStorage.getItem("authToken");
        }

        function getAuthHeaders() {
            const token = localStorage.getItem("authToken");
            const h = {};
            if (token) h["Authorization"] = `Bearer ${token}`;
            return h;
        }

        const moduleStartedWithToken = hasAuthToken();

        function handleAuthError() {
            // limpia información de sesión y recarga para mostrar login
            localStorage.removeItem("authToken");
            localStorage.removeItem("loggedUser");
            localStorage.removeItem("userId");
            localStorage.removeItem("avatarUrl");

            // Si antes teníamos token, avisamos y recargamos una vez (forzamos login).
            if (moduleStartedWithToken) {
                alert("La sesión ha expirado o no está autorizada. Serás redirigido al login.");
                window.location.reload();
            } else {
                // no había token al arrancar: sólo renderizamos vacío (sin alert ni reload)
                TASKS = [];
                applyFilters();
                // renderTasks(TASKS);
            }
        }

        function renderTasks(tasks) {
            const container = document.getElementById("task-list");
            container.innerHTML = "";

            function getCategoryNameById(catId) {
                if (catId == null || catId === "") return "";
                const cats = Array.isArray(window.cats) ? window.cats : [];
                const found = cats.find(c => String(c.id) === String(catId));
                return found ? (found.name || "") : "";
            }

            if (!tasks || !tasks.length) {
                container.innerHTML = "<p>No hay tareas para mostrar.</p>";
                return;
            }

            tasks.forEach(t => {
                const statusClass = t.status === "no_started" ? "task-no" :
                                    t.status === "started" ? "task-started" :
                                    "task-ended";

                const taskEl = document.createElement("div");
                taskEl.className = `task-item ${statusClass}`;

                // date box
                const dateBox = document.createElement("div");
                dateBox.className = "task-date";
                dateBox.textContent = formatDate(t.due_date);

                // body (ahora layout horizontal: category | text)
                const body = document.createElement("div");
                body.className = "task-body";

                // category pill (only if exists) - se coloca antes del texto
                const categoryName = getCategoryNameById(t.category_id);
                if (categoryName) {
                    const catPill = document.createElement("div");
                    catPill.className = "task-category";
                    catPill.textContent = categoryName;
                    body.appendChild(catPill);
                }

                // task text (ocupa el resto del espacio)
                const textEl = document.createElement("div");
                textEl.className = "task-text";
                textEl.textContent = t.text;
                body.appendChild(textEl);

                // menu (three bars)
                const menu = document.createElement("div");
                menu.className = "task-menu";
                menu.innerHTML = "≡";
                menu.title = "Opciones";

                // compose
                taskEl.appendChild(dateBox);
                taskEl.appendChild(body);
                taskEl.appendChild(menu);

                // data attribute for future actions
                taskEl.dataset.taskId = t.id;

                container.appendChild(taskEl);
            });
        }

        // Carga tareas del backend (GET /tareas/usuario)
        async function fetchTasksFromApi() {
            const container = document.getElementById("task-list");
            container.innerHTML = "<p>Loading tasks...</p>";

            try {
                const res = await fetch(`${API_BASE}/usuario`, {
                    method: "GET",
                    headers: {
                        "Accept": "application/json",
                        ...getAuthHeaders()
                    }
                });

                if (res.status === 401) {
                    handleAuthError();
                    return;
                }

                if (!res.ok) {
                    throw new Error(`GET tareas failed: ${res.status}`);
                }

                const json = await res.json();
                // la API puede devolver { tareas: [...] } o un array directamente
                let arr = [];
                if (Array.isArray(json)) arr = json;
                else if (Array.isArray(json.tareas)) arr = json.tareas;
                else if (Array.isArray(json.tasks)) arr = json.tasks;
                else arr = [];

                arr = arr.map(t => ({
                    ...t,
                    status: mapFromBackendStatus(t.status)
                }));

                // Orden simple por created_at descendente (opcional)
                arr.sort((a,b) => {
                    const ta = new Date(a.created_at || 0).getTime();
                    const tb = new Date(b.created_at || 0).getTime();
                    return tb - ta;
                });

                TASKS = arr;
                applyFilters();
                // renderTasks(TASKS);
            } catch (err) {
                console.error(err);
                document.getElementById("task-list").innerHTML = "<p>Error loading tasks.</p>";
            }
        }

        // POST nueva tarea -> /tareas
        async function createTaskApi({ text, category_id, due_date }) {
            try {
                const res = await fetch(API_BASE, {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                        ...getAuthHeaders()
                    },
                    body: JSON.stringify({ text, category_id, due_date })
                });

                if (res.status === 401) {
                    handleAuthError();
                    return null;
                }

                if (!res.ok) {
                    let bodyText = await res.text().catch(()=>"");
                    try {
                        const parsed = JSON.parse(bodyText);
                        throw new Error(parsed.error || parsed.message || JSON.stringify(parsed));
                    } catch (e) {
                        throw new Error(bodyText || `POST failed: ${res.status}`);
                    }
                }

                const created = await res.json();
                if (created && created.status) created.status = mapFromBackendStatus(created.status);
                return created;
            } catch (err) {
                console.error("Create task error:", err);
                alert("No se pudo crear la tarea: " + err.message);
                return null;
            }
        }

        // PUT actualizar tarea -> /tareas/:id
        async function updateTaskApi(id, payload) {
            try {
                const res = await fetch(`${API_BASE}/${id}`, {
                    method: "PUT",
                    headers: {
                        "Content-Type": "application/json",
                        ...getAuthHeaders()
                    },
                    body: JSON.stringify(payload)
                });

                if (res.status === 401) {
                    handleAuthError();
                    return null;
                }

                if (!res.ok) {
                    const txt = await res.text().catch(()=>"");
                    throw new Error(txt || `PUT failed: ${res.status}`);
                }

                const updated = await res.json();
                if (updated && updated.status) updated.status = mapFromBackendStatus(updated.status);
                return updated;
            } catch (err) {
                console.error("Update task error:", err);
                alert("No se pudo actualizar la tarea: " + err.message);
                return null;
            }
        }

        // DELETE tarea -> /tareas/:id
        async function deleteTaskApi(id) {
            try {
                const res = await fetch(`${API_BASE}/${id}`, {
                    method: "DELETE",
                    headers: {
                        ...getAuthHeaders()
                    }
                });

                if (res.status === 401) {
                    handleAuthError();
                    return false;
                }

                if (!res.ok) {
                    const txt = await res.text().catch(()=>"");
                    throw new Error(txt || `DELETE failed: ${res.status}`);
                }

                // respuesta esperada: { id: .., message: "..." }
                await res.json().catch(()=>null);
                return true;
            } catch (err) {
                console.error("Delete task error:", err);
                alert("No se pudo eliminar la tarea: " + err.message);
                return false;
            }
        }

        // inicializar: pedir tareas desde API
        // fetchTasksFromApi();

        // sólo pedir tareas si hay token (evitamos 401 en carga anónima)
        if (hasAuthToken()) {
            fetchTasksFromApi();
        } else {
            // limpiar estado y mostrar nada hasta que haya login
            TASKS = [];
            applyFilters();
            // renderTasks(TASKS);
            // opcional: preparar select de categorías vacío
            if (selectCategory) selectCategory.innerHTML = `<option value="">— No category —</option>`;
        }

        // Cuando haya un login, recargar tareas (usará el token actualizado en localStorage)
        window.addEventListener('app:login', (ev) => {
            fetchTasksFromApi();
        });

        window.addEventListener('app:logout', () => {
            TASKS = [];
            applyFilters();
            // renderTasks(TASKS);
            // limpiar select de categorías también para evitar valores fantasma
            if (selectCategory) {
                selectCategory.innerHTML = `<option value="">— No category —</option>`;
            }
        });

        // re-render cuando las categorías estén listas (si usas el patrón event)
        window.addEventListener('categories:loaded', () => {
            // renderTasks(TASKS);
            applyFilters();
            populateCategorySelect(); // actualiza el select de crear tarea
        });

        // rellenar select de categorias leyendo el DOM de categorías
        function populateCategorySelect() {
            selectCategory.innerHTML = `<option value="">— No category —</option>`;
            const catLabels = document.querySelectorAll("#category-list .category label");
            if (catLabels && catLabels.length) {
                catLabels.forEach(lbl => {
                    const forAttr = lbl.getAttribute("for");
                    let id = "";
                    if (forAttr && forAttr.startsWith("cat-")) id = forAttr.replace("cat-","");
                    const opt = document.createElement("option");
                    opt.value = id || "";
                    opt.textContent = lbl.textContent.trim();
                    selectCategory.appendChild(opt);
                });
            } else {
                // si no hay DOM (aún no cargadas categorías), intenta usar window.cats
                const cats = Array.isArray(window.cats) ? window.cats : [];
                cats.forEach(c => {
                    const opt = document.createElement("option");
                    opt.value = c.id;
                    opt.textContent = c.name || `#${c.id}`;
                    selectCategory.appendChild(opt);
                });
            }
        }

        const addBar = document.getElementById("add-task-bar");
        if (addBar) {
            addBar.addEventListener("click", () => {
                populateCategorySelect();
                // reset inputs
                txtTask.value = "";
                inputDue.value = "";
                selectCategory.value = "";
                modalTask.style.display = "block";
            });
        }

        modalTaskClose.onclick = () => modalTask.style.display = "none";
        window.addEventListener("click", (e) => {
            if (e.target === modalTask) modalTask.style.display = "none";
        });

        btnSaveTask.addEventListener("click", async () => {
            const text = txtTask.value.trim();
            const due = inputDue.value;
            const cat = selectCategory.value;

            if (!text) {
                alert("Please enter a description for the task.");
                txtTask.focus();
                return;
            }
            if (!due) {
                if (!confirm("No due date provided. Create task without due date?")) return;
            }

            // prepara datos para API
            const payload = {
                text: text
            };

            if (cat !== "" && cat != null) {
                const n = Number(cat);
                if (!Number.isNaN(n)) payload.category_id = n;
            }

            if (due) {
                payload.due_date = new Date(due).toISOString();
            }

            const created = await createTaskApi(payload);
            if (created) {
                TASKS.unshift(created);
                applyFilters();
                // renderTasks(TASKS);

                const list = document.getElementById("task-list");
                const first = list.firstElementChild;
            }

            modalTask.style.display = "none";
        });

        function snapshotFromInputs() {
            return {
                text: (editText.value || "").trim(),
                due: editDue.value || "",
                status: editStatus.value || ""
            };
        }

        function isChanged() {
            const s = snapshotFromInputs();
            return JSON.stringify(s) !== JSON.stringify(originalEditSnapshot);
        }

        function setSaveEnabled(enabled) {
            btnSave.disabled = !enabled;
        }

        const taskListEl = document.getElementById("task-list");
        taskListEl.addEventListener("click", (e) => {
            // si el objetivo es el .task-menu, o contiene el icono, subimos al .task-item
            const menu = e.target.closest(".task-menu");
            if (!menu) return;
            const taskItem = menu.closest(".task-item");
            if (!taskItem) return;
            const id = taskItem.dataset.taskId;
            if (!id) return;

            // buscar la tarea en TASKS
            const task = TASKS.find(x => String(x.id) === String(id));
            if (!task) return;

            // rellenar modal
            currentEditTaskId = task.id;
            editText.value = task.text || "";
            editDue.value = task.due_date ? isoToLocalDatetime(task.due_date) : "";
            editStatus.value = task.status || "no_started";

            // guardar snapshot original
            originalEditSnapshot = snapshotFromInputs();
            setSaveEnabled(false);

            // mostrar modal
            editModal.style.display = "block";
        });

        editClose.onclick = () => { editModal.style.display = "none"; currentEditTaskId = null; };
        window.addEventListener("click", (e) => {
            if (e.target === editModal) {
                editModal.style.display = "none"; currentEditTaskId = null;
            }
        });

        [editText, editDue, editStatus].forEach(inp => {
            inp.addEventListener("input", () => {
                setSaveEnabled(isChanged());
            });
        });

        btnDelete.addEventListener("click", async () => {
            if (currentEditTaskId == null) return;
            const ok = confirm("Are you sure you want to delete this task? This action cannot be undone.");
            if (!ok) return;

            const success = await deleteTaskApi(currentEditTaskId);
            if (success) {
                const idx = TASKS.findIndex(x => String(x.id) === String(currentEditTaskId));
                if (idx >= 0) {
                    TASKS.splice(idx, 1);
                    applyFilters();
                    //renderTasks(TASKS);
                }
            }

            editModal.style.display = "none";
            currentEditTaskId = null;
        });

        btnSave.addEventListener("click", async () => {
            if (currentEditTaskId == null) return;
            const idx = TASKS.findIndex(x => String(x.id) === String(currentEditTaskId));
            if (idx < 0) return;

            // validar: descripción no vacía
            const newText = editText.value.trim();
            if (!newText) {
                alert("Please enter a description.");
                editText.focus();
                return;
            }

            const payload = {
                status: mapToBackendStatus(editStatus.value),
                text: newText,
                due_date: editDue.value ? new Date(editDue.value).toISOString() : null
            };

            Object.keys(payload).forEach(k => payload[k] === undefined && delete payload[k]);

            const updated = await updateTaskApi(currentEditTaskId, payload);
            if (updated) {
                // reemplazar en TASKS
                TASKS[idx] = updated;
                applyFilters();
                // renderTasks(TASKS);

                // scrollear a la tarea actualizada
                const newTaskEl = document.querySelector(`.task-item[data-task-id="${TASKS[idx].id}"]`);
                if (newTaskEl) newTaskEl.scrollIntoView({ behavior: "smooth", block: "center" });
            }

            editModal.style.display = "none";
            currentEditTaskId = null;
        });

    })
    .catch(err => console.error("Error loading tasks template:", err));