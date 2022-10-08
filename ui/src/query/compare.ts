import { useQuery } from 'react-query';
import Api from './api';

export namespace DiffData {
  export type DiffAddedItem = {
    breaking: boolean;
    method: string;
    path: string;
    message: string;
    description: string;
  };

  export type DiffDeletedItem = DiffAddedItem;

  export type SecurityDetailItem = {
    name: string;
    action: string;
    breaking: boolean;
    message: string;
  };

  export type DiffModifiedItemSecurity = {
    details: SecurityDetailItem[];
    message: string;
  };

  export type ParametersDetailItem = {
    name: string;
    action: string;
    breaking: boolean;
    deprecated: boolean;
    description: string;
    in: string;
    message: string;
  };

  export type DiffModifiedItemParameters = {
    breaking: boolean;
    details: ParametersDetailItem[];
    message: string;
  };

  export type DiffModifiedItemNew = {
    description: string;
    operationId: string;
    produces: string[];
    responses: unknown;
  };

  export type DiffModifiedItem = DiffAddedItem & {
    breaking: boolean;
    summary: string;
    description: string;
    requestBody: string;
    security: DiffModifiedItemSecurity;
    parameters: DiffModifiedItemParameters;
    new: unknown;
    old: unknown;
  };

  export type JsonDiffResult = {
    breaking: boolean;
    deprecated: string;
    added: DiffAddedItem[];
    modified: DiffModifiedItem[];
    deleted: DiffDeletedItem[];
  };

  type DiffInfo = {
    id: string;
    new_spec_id: string;
    old_spec_id: string;
    service_id: string;
    status: string;
    created_at: string;
    updated_at: string;
  };
  export type JsonDiff = DiffInfo & {
    result: {
      json: JsonDiffResult;
    };
  };

  export type MarkdownDiff = DiffInfo & {
    config: {
      output_format: string;
    };
    result: {
      markdown: string;
    };
  };
}

export function useFetchCompare(
  serviceId: string,
  spec1Id: string,
  spec2Id: string,
) {
  const url = `/services/${serviceId}/specs/diff`;
  const payload = {
    old_spec_id: spec1Id,
    new_spec_id: spec2Id,
  };

  const cfg = {
    enabled: false,
  };

  return useQuery(
    ['compare', serviceId, spec1Id, spec2Id],
    async () => {
      const result = await Api.post(url, payload);
      return result;
    },
    cfg,
  );
}

export function useFetchMarkdown(
  serviceId: string,
  spec1Id: string,
  spec2Id: string,
) {
  const url = `/services/${serviceId}/specs/diff`;
  const payload = {
    old_spec_id: spec1Id,
    new_spec_id: spec2Id,
    config: { output_format: 'markdown' },
  };

  const cfg = {
    enabled: false,
  };

  return useQuery(
    ['compare', serviceId, spec1Id, spec2Id, 'markdown'],
    async () => {
      const result = await Api.post(url, payload);
      return result;
    },
    cfg,
  );
}
