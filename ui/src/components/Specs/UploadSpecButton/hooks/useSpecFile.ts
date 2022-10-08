import { useState } from 'react';
import YAML from 'yaml';

export type SpecFileData = {
  text: string;
  parsedSpec?: {
    openapi: string;
    swagger: string;
  };
};

type ReadSpecFunciton = (file: File) => Promise<SpecFileData>;

export default function useSpecFile() {
  const [reading, setReading] = useState(false);
  const [error, setError] = useState(null);

  const readSpec = (file: File) => {
    setReading(true);
    return new Promise((resolve: (value: SpecFileData) => void) => {
      const reader = new FileReader();
      reader.onloadend = () => {
        try {
          const text = reader.result as string;
          const parsed = file.name.includes('.json')
            ? JSON.parse(text)
            : YAML.parse(text);
          setReading(false);
          resolve({ text, parsedSpec: parsed });
        } catch (err) {
          setReading(false);
          setError(err);
          resolve({ text: '' });
        }
      };
      reader.readAsText(file);
    });
  };

  return [readSpec as ReadSpecFunciton, reading, error, setError];
}
