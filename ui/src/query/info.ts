import { useQuery } from 'react-query';
import Api from './api';

export namespace InfoData {
  export type Info = {
    auth?: {
      enabled: boolean;
    };
  };
}

export function useFetchInfo() {
  return useQuery('info', () => Api.get('/info'));
}
