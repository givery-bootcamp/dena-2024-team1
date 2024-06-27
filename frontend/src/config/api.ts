const API_URL = import.meta.env.VITE_API_ENDPOINT_PATH || "http://localhost:9000";

// axios config
export const config = {
  baseURL: API_URL,
  withCredentials: true,
};
