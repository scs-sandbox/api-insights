import { useQuery, useMutation } from 'react-query';
import Api from './api';

export namespace SpecData {
  export enum SpecState {
    Archive = 'Archive',
    Release = 'Release',
    Development = 'Development',
    Latest = 'Latest',
    Reconstructed = 'Reconstructed',
  }

  export type Spec = {
    created_at: string;
    updated_at: string;
    id: string;
    service_id: string;
    version: string;
    revision: string;
    doc: string;
    doc_type: string;
    score: number;
    state: SpecState;
    valid: string;
  };

  export type Snapshot = {
    created_at: string;
    updated_at: string;
    id: string;
    service_id: string;
    name: string;
    doc: string;
    doc_type: string;
  };
}

type AddSpecPayload = {
  service_id: string;
  revision: string;
  doc: string;
};

export function useFetchSpecList(serviceId: string) {
  return useQuery(['spec-list', serviceId], () => {
    const url = `/services/${serviceId}/specs`;
    if (!serviceId) return [];
    return Api.get(url);
  });
}

export function useFetchSpecDetail(serviceId: string, specId: string) {
  return useQuery(['spec-detail', serviceId], () => {
    const url = `/services/${serviceId}/specs/${specId}?withDoc=true`;
    if (!serviceId || !specId) return null;
    return Api.get(url);
  });
}

export function useAddSpec() {
  return useMutation((payload: AddSpecPayload) => {
    const url = `/services/${payload.service_id}/specs`;

    return Api.post(url, payload);
  });
}

export function useReconstructSnapshot(serviceId: string) {
  return useMutation(() => {
    const url = `/services/${serviceId}/specs/reconstruct`;

    return Api.post(url, {});
  });
}

export function useReleaseSpec() {
  return useMutation((id: string) => {
    const url = `/specs/${id}`;

    return Api.patch(url);
  });
}

export function useArchiveSpec() {
  return useMutation((id: string) => {
    const url = `/specs/${id}`;

    return Api.patch(url);
  });
}
