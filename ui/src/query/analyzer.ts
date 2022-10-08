import { useQuery } from 'react-query';
import Api from './api';

export namespace AnalyserData {
  export type Analyser = {
    status: string;
    name_id: string;
    title: string;
  };
}

export function useFetchAnalyzerList() {
  return useQuery('analyzer-list', () => Api.get('/analyzers?status=active'));
}
export async function TriggerReanalyze(
  serviceId: string,
  specId: string,
  analyzers: string[],
) {
  const url = `/services/${serviceId}/specs/${specId}/analyses`;
  const payload = {
    analyzers,
  };
  const result = await Api.post(url, payload);
  return result;
}
