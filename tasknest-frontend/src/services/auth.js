class AuthService {
  constructor() {
    this.token = localStorage.getItem("auth_token");
  }

  async login(email, password) {
    const response = await fetch("http://localhost:8080/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });

    if (!response.ok) {
      throw new Error("Login failed");
    }

    const data = await response.json();
    this.token = data.token;
    localStorage.setItem("auth_token", this.token);
    return data;
  }

  logout() {
    this.token = null;
    localStorage.removeItem("auth_token");
  }

  getAuthHeaders() {
    return this.token ? { Authorization: `Bearer ${this.token}` } : {};
  }

  isAuthenticated() {
    return !!this.token;
  }
}

export default new AuthService();
