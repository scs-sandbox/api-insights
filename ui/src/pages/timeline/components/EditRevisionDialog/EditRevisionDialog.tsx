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

import { useState, ChangeEvent } from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogBody from '../../../../components/Dialog/DialogBody/DialogBody';
import FieldItem from '../../../../components/FieldItem/FieldItem';
import DialogActions from '../../../../components/Dialog/DialogActions/DialogActions';
import DialogTitle from '../../../../components/Dialog/DialogTitle/DialogTitle';
import IconButton from '../../../../components/IconButton/IconButton';
import BusyIcon from '../../../../components/BusyIcon/BusyIcon';
import { ServiceData } from '../../../../query/service';
// import './EditGroupDialog.scss';

export type EditingEvent = {
  id: string;
  title: string;
  product_tag: string;
};

type Props = {
  service: ServiceData.Service;
  open: boolean;
  busy: boolean;
  handleClose?: () => void;
  onEditing: (e: EditingEvent) => void;
};

export default function EditGroupDialog(props: Props) {
  const [title, setTitle] = useState(props.service.title);
  const [tag, setTag] = useState(props.service.product_tag);

  const trimmedTitle = title.trim();
  const trimmedTag = tag.trim();

  const onChangeTitle = (e: ChangeEvent<HTMLInputElement>) => {
    setTitle(e.target.value);
  };

  const onChangeTag = (e: ChangeEvent<HTMLInputElement>) => {
    setTag(e.target.value);
  };

  const onEditing = () => {
    if (props.onEditing) {
      props.onEditing({
        id: props.service?.id,
        title: trimmedTitle,
        product_tag: trimmedTag,
      });
    }
  };

  const invalidInputs = !trimmedTitle || !trimmedTag;
  const onClose = props.busy ? undefined : props.handleClose;

  return (
    <Dialog open={props.open} fullWidth maxWidth="sm">
      <DialogTitle onClose={onClose}>Edit Revision</DialogTitle>
      <DialogBody className="edit-group-dialog-body light-bg">
        <FieldItem label="What is this service called?">
          <input value={title} onChange={onChangeTitle} />
        </FieldItem>
        <FieldItem label="Assign product affiliation">
          <input value={tag} onChange={onChangeTag} />
        </FieldItem>
      </DialogBody>
      <DialogActions className="edit-group-dialog-actions">
        <Button onClick={onClose}>Cancel</Button>
        <IconButton
          icon={props.busy && <BusyIcon busy />}
          disabled={invalidInputs || props.busy}
          onClick={onEditing}
        >
          Update
        </IconButton>
      </DialogActions>
    </Dialog>
  );
}
