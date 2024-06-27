import axios from "axios";

import { Configuration } from "~/generated/configuration";

export const API_ENDPOINT_PATH = import.meta.env.VITE_API_ENDPOINT_PATH || "http://localhost:9000";

export const config = new Configuration({
  basePath: API_ENDPOINT_PATH,
});

export const axiosInstance = axios.create({
  baseURL: API_ENDPOINT_PATH,
  withCredentials: true,
});
