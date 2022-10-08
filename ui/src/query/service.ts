import { useQuery, useMutation } from 'react-query';
import Api from './api';

export namespace ServiceData {
  export type ServiceContact = {
    name: string;
    url: string;
    email: string;
  };

  export type ServiceSummary = {
    score: number;
    version: string;
    revision: string;
    updated_at: string;
  };

  export type CreateServiceData = {
    name_id: string;
    title: string;
    description: string;
    contact?: ServiceContact;
    organization_id: string;
    product_tag: string;
    summary?: ServiceSummary;
  };

  export type Service = {
    visibility: string;
    created_at: string;
    updated_at: string;
    id: string;
    name_id: string;
    title: string;
    description: string;
    contact?: ServiceContact;
    organization_id: string;
    product_tag: string;
    summary?: ServiceSummary;
  };
}

export type AddServicePayload = {
  contact: ServiceData.ServiceContact;
  description: string;
  name_id: string;
  organization_id: string;
  product_tag: string;
  title: string;
};

type PatchServicePayload = {
  contact?: ServiceData.ServiceContact;
  description?: string;
  name_id?: string;
  organization_id?: string;
  product_tag?: string;
  title?: string;
};

export type PatchServiceData = PatchServicePayload & {
  id: string;
};

export function useFetchServiceList() {
  return useQuery('service-list', () => Api.get('/services'));
}

export function useAddService() {
  return useMutation((payload: AddServicePayload) => {
    const url = '/services';
    return Api.post(url, payload);
  });
}

export function usePatchService() {
  return useMutation((payload: PatchServiceData) => {
    const { id, ...other } = payload;
    const url = `/services/${id}`;

    return Api.patch(url, other as PatchServicePayload);
  });
}
