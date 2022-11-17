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

import { ChangeEventHandler, useState, ChangeEvent } from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import FileInput from '../../FileInput/FileInput';
import DialogBody from '../../Dialog/DialogBody/DialogBody';
import FieldItem from '../../FieldItem/FieldItem';
import DialogActions from '../../Dialog/DialogActions/DialogActions';
import DialogTitle from '../../Dialog/DialogTitle/DialogTitle';
import BusyIcon from '../../BusyIcon/BusyIcon';
import IconButton from '../../IconButton/IconButton';
import './UploadSpecDialog.scss';

export type UploadingEvent = {
  file: File;
  revision: string;
};

type Props = {
  open: boolean;
  busy: boolean;
  handleClose?: () => void;
  onUploading?: (e: UploadingEvent) => void;
};

export default function UploadSpecDialog(props: Props) {
  const [revision, setRevision] = useState('');
  const [file, setFile] = useState<File>();

  const trimmedRevision = revision.trim();

  const onChangeRevision = (e: ChangeEvent<HTMLInputElement>) => {
    setRevision(e.target.value);
  };

  const onFileSelected: ChangeEventHandler<HTMLInputElement> = (e) => {
    const selectedFile = e.target.files ? e.target.files[0] : undefined;
    setFile(selectedFile);
  };

  const onUploading = () => {
    if (!file) return;

    const e: UploadingEvent = {
      file,
      revision: trimmedRevision,
    };

    if (props.onUploading) {
      props.onUploading(e);
    }
  };

  const invalidInputs = !trimmedRevision || !file;
  const fileName = file ? file.name : '';
  const onClose = props.busy ? undefined : props.handleClose;

  return (
    <Dialog open={props.open} fullWidth maxWidth="md">
      <DialogTitle onClose={onClose}>Upload New Spec</DialogTitle>
      <DialogBody className="upload-spec-dialog-body">
        <div className="field-group light-bg">
          <FieldItem label="Version">
            Spec version will be detected from the uploaded spec file.
          </FieldItem>
          <FieldItem>
            <div className="upload-spec-field">
              <div className="upload-spec-title">
                Select spec file to upload
              </div>
              <FileInput
                value={fileName}
                accept=".json, .yml, yaml, YAML"
                onChange={onFileSelected}
              />
            </div>
          </FieldItem>
        </div>
        <div className="field-group">
          <FieldItem label="Enter revision version">
            <div>This is an identifier when iterating on the same version.</div>
          </FieldItem>
          <FieldItem>
            <input
              placeholder="Name your revision, eg Beta-1"
              value={revision}
              onChange={onChangeRevision}
            />
          </FieldItem>
        </div>
      </DialogBody>
      <DialogActions>
        <Button disabled={props.busy} onClick={onClose}>
          Cancel
        </Button>
        <IconButton
          icon={props.busy && <BusyIcon busy />}
          disabled={invalidInputs || props.busy}
          onClick={onUploading}
        >
          Upload
        </IconButton>
      </DialogActions>
    </Dialog>
  );
}
