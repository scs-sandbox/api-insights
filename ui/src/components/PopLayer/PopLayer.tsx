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

import { useEffect, useState, HTMLAttributes } from 'react';
import { createPortal } from 'react-dom';
import './PopLayer.scss';

type Props = HTMLAttributes<HTMLElement>;

export default function PopLayer(props: Props) {
  const [popLayer, setPopLayer] = useState(
    document.querySelector('#pop-layer'),
  );

  useEffect(() => {
    if (popLayer) return;

    const domPopLayer = document.createElement('div');
    domPopLayer.id = 'pop-layer';
    document.body.appendChild(domPopLayer);

    setPopLayer(domPopLayer);
  }, []);

  return popLayer ? createPortal(props.children, popLayer) : null;
}
