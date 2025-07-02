import AuthService from "./auth.js";

const API_BASE_URL = "http://localhost:8080"; // Your Go backend URL

class TaskAPI {
  async request(endpoint, options = {}) {
    const url = `${API_BASE_URL}${endpoint}`;
    const config = {
      headers: {
        "Content-Type": "application/json",
        ...AuthService.getAuthHeaders(), // Add auth headers
        ...options.headers,
      },
      ...options,
    };

    try {
      const response = await fetch(url, config);

      if (response.status === 401) {
        // Token expired or invalid
        AuthService.logout();
        window.location.href = "/login";
        return;
      }

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      // Handle empty responses (like DELETE)
      const text = await response.text();
      return text ? JSON.parse(text) : null;
    } catch (error) {
      console.error("API request failed:", error);
      throw error;
    }
  }

  // Get all tasks
  async getTasks() {
    return this.request("/tasks");
  }

  // Create a new task
  async createTask(task) {
    return this.request("/tasks", {
      method: "POST",
      body: JSON.stringify({
        title: task.title,
        description: task.description,
        status: task.status,
        completed: task.status === "completed",
      }),
    });
  }

  // Update a task
  async updateTask(id, task) {
    return this.request(`/tasks/${id}`, {
      method: "PUT",
      body: JSON.stringify({
        title: task.title,
        description: task.description,
        status: task.status,
        completed: task.status === "completed",
      }),
    });
  }

  // Delete a task
  async deleteTask(id) {
    return this.request(`/tasks/${id}`, {
      method: "DELETE",
    });
  }
}

export default new TaskAPI();
