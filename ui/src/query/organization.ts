import { useQuery } from 'react-query';
import Api from './api';

export namespace OrganizationData {
  export type Organization = {
    id: string;
    name_id: string;
    title: string;
    description: string;
  };
}

export function useFetchOrganizationList() {
  return useQuery('organization-list', () => Api.get('/organizations'));
}
