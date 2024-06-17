import axios, { AxiosResponse , AxiosInstance } from "axios";
import { mapKeys, camelCase, mapValues, isArray, isObject } from "lodash";

type ApiResponse<T> = {
  data: T
} & AxiosResponse

const API_ENDPOINT_PATH: string = import.meta.env.VITE_API_ENDPOINT_PATH ?? "";
export const client: AxiosInstance = axios.create({ baseURL: API_ENDPOINT_PATH });

const mapKeysDeep = <T>(data: T, callback: (value: unknown, key: string) => string): T => {
  if (isArray(data)) {
    return data.map(innerData => mapKeysDeep(innerData, callback)) as T;
  } else if (isObject(data)) {
    return mapValues(mapKeys(data, callback), val =>
      mapKeysDeep(val, callback),
    ) as T;
  } else {
    return data;
  }
};

const mapKeysCamelCase = <T>(data: T): T =>
  mapKeysDeep(data, (_: unknown, key: string) => camelCase(key));

client.interceptors.response.use(
  (response: AxiosResponse<unknown, unknown>) => {
    const { data } = response;
    const convertedData = mapKeysCamelCase(data);
    return { ...response, data: convertedData } as ApiResponse<unknown>;
  },
  (error) => {
    console.log(error);
    return Promise.reject(error);
  },
);
