<template>
  <div class="app-container">
    <!-- Background -->
    <div class="background-image"></div>
    <div class="background-overlay"></div>

    <!-- Loading overlay -->
    <div v-if="loading" class="loading-overlay">
      <div class="loading-card">
        <div class="spinner"></div>
        <p>Loading tasks...</p>
      </div>
    </div>

    <!-- Main content -->
    <div class="main-content">
      <!-- Header -->
      <header class="app-header">
        <div class="header-left">
          <div class="logo">
            <svg
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
            >
              <path d="M9 11l3 3 8-8" />
              <path
                d="M21 12c0 4.97-4.03 9-9 9s-9-4.03-9-9 4.03-9 9-9c1.51 0 2.93.37 4.18 1.03"
              />
            </svg>
          </div>
          <h1>TaskNest</h1>
        </div>
        <div class="task-count">{{ tasks.length }} tasks</div>
      </header>

      <!-- Error message -->
      <div v-if="error" class="error-card">
        <p>{{ error }}</p>
        <button @click="fetchTasks" class="retry-btn">Try Again</button>
      </div>

      <!-- Controls -->
      <div class="controls">
        <div class="control-group">
          <select v-model="filterStatus" class="glass-select">
            <option value="all">Filter by status...</option>
            <option value="pending">Pending</option>
            <option value="completed">Completed</option>
            <option value="archived">Archived</option>
          </select>
        </div>
        <div class="control-group">
          <select v-model="sortBy" class="glass-select">
            <option value="created">Sort by...</option>
            <option value="created">Date Created</option>
            <option value="title">Title</option>
            <option value="status">Status</option>
          </select>
        </div>
      </div>

      <!-- Tasks List -->
      <div class="tasks-container">
        <div v-if="filteredTasks.length === 0 && !loading" class="empty-state">
          <div class="empty-icon">
            <svg
              width="48"
              height="48"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="1.5"
            >
              <path d="M9 11l3 3 8-8" />
              <path
                d="M21 12c0 4.97-4.03 9-9 9s-9-4.03-9-9 4.03-9 9-9c1.51 0 2.93.37 4.18 1.03"
              />
            </svg>
          </div>
          <h3>No tasks found</h3>
          <p>Create your first task to get started!</p>
        </div>

        <div v-for="task in filteredTasks" :key="task.id" class="task-card">
          <div class="task-content">
            <h3 class="task-title">{{ task.title }}</h3>
            <p class="task-description">{{ task.description }}</p>
            <p class="task-date" v-if="task.created_at">
              Created: {{ formatDate(task.created_at) }}
            </p>
          </div>
          <div class="task-actions">
            <span :class="['status-badge', `status-${task.status}`]">
              {{ task.status }}
            </span>
            <button
              @click="toggleTaskStatus(task)"
              :disabled="updating"
              class="action-btn"
            >
              <svg
                width="16"
                height="16"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <polyline points="20,6 9,17 4,12" />
              </svg>
            </button>
            <button
              @click="deleteTask(task.id)"
              :disabled="updating"
              class="action-btn delete-btn"
            >
              <svg
                width="16"
                height="16"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <polyline points="3,6 5,6 21,6" />
                <path
                  d="M19,6v14a2,2 0 0,1-2,2H7a2,2 0 0,1-2-2V6m3,0V4a2,2 0 0,1,2-2h4a2,2 0 0,1,2,2v2"
                />
              </svg>
            </button>
          </div>
        </div>
      </div>

      <!-- Floating Add Button -->
      <button @click="showModal = true" class="fab">
        <svg
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <line x1="12" y1="5" x2="12" y2="19" />
          <line x1="5" y1="12" x2="19" y2="12" />
        </svg>
      </button>

      <!-- Modal -->
      <div v-if="showModal" class="modal-overlay" @click="closeModal">
        <div class="modal-content" @click.stop>
          <h2>Create Task</h2>

          <div class="form-group">
            <label>Title</label>
            <input
              v-model="newTask.title"
              type="text"
              placeholder="Enter task title"
              class="glass-input"
            />
          </div>

          <div class="form-group">
            <label>Description</label>
            <textarea
              v-model="newTask.description"
              placeholder="Enter task description"
              rows="3"
              class="glass-input"
            ></textarea>
          </div>

          <div class="form-group">
            <label>Status</label>
            <select v-model="newTask.status" class="glass-input">
              <option value="pending">Pending</option>
              <option value="completed">Completed</option>
              <option value="archived">Archived</option>
            </select>
          </div>

          <div class="modal-actions">
            <button
              @click="createTask"
              :disabled="creating || !newTask.title.trim()"
              class="btn-primary"
            >
              {{ creating ? "Creating..." : "Save" }}
            </button>
            <button
              @click="closeModal"
              :disabled="creating"
              class="btn-secondary"
            >
              Cancel
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import TaskAPI from "../services/api.js";

// Reactive data
const tasks = ref([]);
const loading = ref(false);
const error = ref(null);
const updating = ref(false);
const creating = ref(false);
const showModal = ref(false);
const filterStatus = ref("all");
const sortBy = ref("created");

const newTask = ref({
  title: "",
  description: "",
  status: "pending",
});

// Computed properties
const filteredTasks = computed(() => {
  let filtered = tasks.value;

  if (filterStatus.value !== "all") {
    filtered = filtered.filter((task) => task.status === filterStatus.value);
  }

  filtered.sort((a, b) => {
    switch (sortBy.value) {
      case "title":
        return a.title.localeCompare(b.title);
      case "status":
        return a.status.localeCompare(b.status);
      case "created":
      default:
        return (
          new Date(b.created_at || b.createdAt) -
          new Date(a.created_at || a.createdAt)
        );
    }
  });

  return filtered;
});

// API Methods
const fetchTasks = async () => {
  loading.value = true;
  error.value = null;

  try {
    const response = await TaskAPI.getTasks();
    tasks.value = response || [];
  } catch (err) {
    error.value =
      "Failed to load tasks. Make sure your Go backend is running on http://localhost:8080";
    console.error("Error fetching tasks:", err);
  } finally {
    loading.value = false;
  }
};

const createTask = async () => {
  if (!newTask.value.title.trim()) return;

  creating.value = true;
  error.value = null;

  try {
    const response = await TaskAPI.createTask(newTask.value);
    tasks.value.unshift(response);
    closeModal();
  } catch (err) {
    error.value = "Failed to create task";
    console.error("Error creating task:", err);
  } finally {
    creating.value = false;
  }
};

const toggleTaskStatus = async (task) => {
  const statusOrder = ["pending", "completed", "archived"];
  const currentIndex = statusOrder.indexOf(task.status);
  const nextIndex = (currentIndex + 1) % statusOrder.length;
  const newStatus = statusOrder[nextIndex];

  updating.value = true;
  error.value = null;

  try {
    const updatedTask = { ...task, status: newStatus };
    await TaskAPI.updateTask(task.id, updatedTask);
    task.status = newStatus;
  } catch (err) {
    error.value = "Failed to update task";
    console.error("Error updating task:", err);
  } finally {
    updating.value = false;
  }
};

const deleteTask = async (taskId) => {
  if (!confirm("Are you sure you want to delete this task?")) return;

  updating.value = true;
  error.value = null;

  try {
    await TaskAPI.deleteTask(taskId);
    tasks.value = tasks.value.filter((task) => task.id !== taskId);
  } catch (err) {
    error.value = "Failed to delete task";
    console.error("Error deleting task:", err);
  } finally {
    updating.value = false;
  }
};

// Utility methods
const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString();
};

const closeModal = () => {
  showModal.value = false;
  newTask.value = {
    title: "",
    description: "",
    status: "pending",
  };
};

// Lifecycle
onMounted(() => {
  fetchTasks();

  // Close modal on escape key
  document.addEventListener("keydown", (e) => {
    if (e.key === "Escape" && showModal.value) {
      closeModal();
    }
  });
});
</script>

<style scoped>
/* All the previous styles remain the same, plus these additions: */

.loading-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.3);
  backdrop-filter: blur(4px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.loading-card {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 16px;
  padding: 2rem;
  text-align: center;
  color: white;
}

.spinner {
  width: 32px;
  height: 32px;
  border: 3px solid rgba(255, 255, 255, 0.3);
  border-top: 3px solid white;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 1rem;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

.error-card {
  background: rgba(239, 68, 68, 0.1);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(239, 68, 68, 0.3);
  border-radius: 12px;
  padding: 1rem;
  margin-bottom: 1rem;
  color: rgb(254, 202, 202);
}

.retry-btn {
  background: none;
  border: none;
  color: rgb(254, 202, 202);
  text-decoration: underline;
  cursor: pointer;
  margin-top: 0.5rem;
}

.task-date {
  color: rgba(255, 255, 255, 0.5);
  font-size: 0.75rem;
  margin-top: 0.25rem;
}

/* All other styles from the previous component remain the same */
.app-container {
  position: relative;
  min-height: 100vh;
  width: 100%;
}

.background-image {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-image: url("https://images.unsplash.com/photo-1506905925346-21bda4d32df4?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=2070&q=80");
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  z-index: -2;
}

.background-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    135deg,
    rgba(255, 200, 150, 0.3) 0%,
    rgba(150, 200, 255, 0.2) 50%,
    rgba(200, 150, 255, 0.3) 100%
  );
  z-index: -1;
}

.main-content {
  position: relative;
  z-index: 1;
  padding: 2rem;
  max-width: 800px;
  margin: 0 auto;
}

.app-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.logo {
  width: 40px;
  height: 40px;
  background: rgba(30, 41, 59, 0.9);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
}

.app-header h1 {
  font-size: 2rem;
  font-weight: 700;
  color: white;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
}

.task-count {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 8px;
  padding: 0.5rem 1rem;
  color: white;
  font-size: 0.875rem;
  font-weight: 500;
}

.controls {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
  margin-bottom: 1.5rem;
}

.control-group {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 12px;
  padding: 0.75rem 1rem;
}

.glass-select {
  background: transparent;
  border: none;
  color: white;
  width: 100%;
  outline: none;
  font-size: 0.875rem;
}

.glass-select option {
  background: rgba(30, 41, 59, 0.95);
  color: white;
}

.tasks-container {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  margin-bottom: 6rem;
}

.empty-state {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 16px;
  padding: 3rem 2rem;
  text-align: center;
  color: white;
}

.empty-icon {
  margin-bottom: 1rem;
  opacity: 0.6;
}

.empty-state h3 {
  font-size: 1.25rem;
  margin-bottom: 0.5rem;
}

.empty-state p {
  opacity: 0.7;
  font-size: 0.875rem;
}

.task-card {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 16px;
  padding: 1.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  transition: all 0.3s ease;
}

.task-card:hover {
  background: rgba(255, 255, 255, 0.15);
  transform: translateY(-2px);
}

.task-content {
  flex: 1;
}

.task-title {
  color: white;
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
}

.task-description {
  color: rgba(255, 255, 255, 0.7);
  font-size: 0.875rem;
}

.task-actions {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.status-badge {
  padding: 0.375rem 0.75rem;
  border-radius: 20px;
  font-size: 0.75rem;
  font-weight: 500;
  backdrop-filter: blur(10px);
  border: 1px solid;
}

.status-pending {
  background: rgba(251, 146, 60, 0.3);
  color: rgb(254, 215, 170);
  border-color: rgba(251, 146, 60, 0.3);
}

.status-completed {
  background: rgba(34, 197, 94, 0.3);
  color: rgb(187, 247, 208);
  border-color: rgba(34, 197, 94, 0.3);
}

.status-archived {
  background: rgba(107, 114, 128, 0.3);
  color: rgb(209, 213, 219);
  border-color: rgba(107, 114, 128, 0.3);
}

.action-btn {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  border: none;
  background: rgba(255, 255, 255, 0.1);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s ease;
}

.action-btn:hover {
  background: rgba(255, 255, 255, 0.2);
}

.action-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.delete-btn:hover:not(:disabled) {
  background: rgba(239, 68, 68, 0.3);
}

.fab {
  position: fixed;
  bottom: 2rem;
  right: 2rem;
  width: 64px;
  height: 64px;
  border-radius: 50%;
  border: none;
  background: rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.3);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.fab:hover {
  background: rgba(255, 255, 255, 0.3);
  transform: scale(1.05);
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(4px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  padding: 1rem;
}

.modal-content {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 16px;
  padding: 2rem;
  width: 100%;
  max-width: 400px;
  color: white;
}

.modal-content h2 {
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-size: 0.875rem;
  font-weight: 500;
  color: rgba(255, 255, 255, 0.8);
}

.glass-input {
  width: 100%;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 8px;
  padding: 0.75rem;
  color: white;
  font-size: 0.875rem;
  outline: none;
  transition: border-color 0.2s ease;
}

.glass-input:focus {
  border-color: rgba(255, 255, 255, 0.4);
}

.glass-input::placeholder {
  color: rgba(255, 255, 255, 0.5);
}

.glass-input option {
  background: rgba(30, 41, 59, 0.95);
  color: white;
}

.modal-actions {
  display: flex;
  gap: 0.75rem;
  margin-top: 1.5rem;
}

.btn-primary {
  flex: 1;
  background: rgba(59, 130, 246, 0.8);
  color: white;
  border: none;
  border-radius: 8px;
  padding: 0.75rem 1.5rem;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.btn-primary:hover:not(:disabled) {
  background: rgba(59, 130, 246, 1);
}

.btn-primary:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.btn-secondary {
  flex: 1;
  background: rgba(255, 255, 255, 0.1);
  color: white;
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 8px;
  padding: 0.75rem 1.5rem;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.btn-secondary:hover:not(:disabled) {
  background: rgba(255, 255, 255, 0.2);
}

.btn-secondary:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

@media (max-width: 768px) {
  .main-content {
    padding: 1rem;
  }

  .controls {
    grid-template-columns: 1fr;
  }

  .task-card {
    flex-direction: column;
    align-items: flex-start;
    gap: 1rem;
  }

  .task-actions {
    align-self: flex-end;
  }
}
</style>
