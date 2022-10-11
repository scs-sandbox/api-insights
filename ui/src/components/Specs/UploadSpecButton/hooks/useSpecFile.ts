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
