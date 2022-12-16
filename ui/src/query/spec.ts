/*
 * Copyright 2022 Cisco Systems, Inc. and its affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

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

export function fetchSpecDetail(serviceId: string, specId: string) {
  const url = `/services/${serviceId}/specs/${specId}?withDoc=true`;
  if (!serviceId || !specId) return null;
  return Api.get(url);
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
