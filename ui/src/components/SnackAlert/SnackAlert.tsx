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

import {
  Alert, AlertColor, Portal, Snackbar,
} from '@mui/material';

type Props = {
  severity: AlertColor;
  message: string;
  autoHideDuration?: number;
  onClose: () => void;
};

export default function SnackAlert(props: Props) {
  const autoHideDuration = props.autoHideDuration || (props.severity === 'error' ? 6000 : 3000);

  return (
    <Portal>
      <Snackbar
        open
        anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
        autoHideDuration={autoHideDuration}
        onClose={props.onClose}
      >
        <Alert
          onClose={props.onClose}
          severity={props.severity}
          sx={{ width: '100%' }}
        >
          {props.message}
        </Alert>
      </Snackbar>
    </Portal>
  );
}
