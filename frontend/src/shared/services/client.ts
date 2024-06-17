/* eslint-disable */
import axios from "axios";
import { mapKeys,camelCase,mapValues,isArray,isObject } from "lodash";

const API_ENDPOINT_PATH = import.meta.env.VITE_API_ENDPOINT_PATH ?? "";
export const client = axios.create({ baseURL: API_ENDPOINT_PATH });
const mapKeysDeep = (data: any, callback: any) => {
    if (isArray(data)) {
      return data.map(innerData => mapKeysDeep(innerData, callback));
    } else if (isObject(data)) {
      return mapValues(mapKeys(data, callback), val =>
        mapKeysDeep(val, callback),
      );
    } else {
      return data;
    }
  };
  
  const mapKeysCamelCase = data =>
    mapKeysDeep(data, (_, key: string) => camelCase(key));
  
client.interceptors.response.use(
    response => {
      const { data } = response;
      const convertedData = mapKeysCamelCase(data);
      return { ...response, data: convertedData };
    },
    error => {
      console.log(error);
      return Promise.reject(error);
    },
);