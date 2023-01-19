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

import { useState, ButtonHTMLAttributes } from 'react';
import { ServiceData, usePatchService } from '../../../../query/service';
import IconButton from '../../../../components/IconButton/IconButton';
import SnackAlert from '../../../../components/SnackAlert/SnackAlert';
import EditGroupDialog, {
  EditingEvent,
} from '../EditRevisionDialog/EditRevisionDialog';
import './EditRevisionButton.scss';

type Props = ButtonHTMLAttributes<HTMLButtonElement> & {
  service: ServiceData.Service;
  onServiceUpdated?: (data: ServiceData.Service) => void;
};

export default function EditGroupButton(props: Props) {
  const [openDialog, setOpenDialog] = useState(false);
  const {
    isLoading: isServicePatching,
    isError: isServicePatchingError,
    isSuccess: isServicePatchingSuccess,
    mutate: patchService,
    reset: resetPatchServiceResult,
  } = usePatchService();

  const onOpenDialog = () => {
    setOpenDialog(true);
  };

  const onCloseDialog = () => {
    setOpenDialog(false);
  };

  const onEditingService = (e: EditingEvent) => {
    const notChanged = e.title === props.service.title
      && e.product_tag === props.service.product_tag;

    if (notChanged) {
      setOpenDialog(false);
      return;
    }

    patchService(e, {
      onSuccess: (data: ServiceData.Service) => {
        setOpenDialog(false);
        if (props.onServiceUpdated) {
          props.onServiceUpdated(data);
        }
      },
    });
  };

  const onCloseMessage = () => {
    resetPatchServiceResult();
  };

  const renderDialog = () => {
    if (!openDialog) return null;

    return (
      <EditGroupDialog
        service={props.service}
        busy={isServicePatching}
        open
        handleClose={onCloseDialog}
        onEditing={onEditingService}
      />
    );
  };

  const renderErrorMessage = () => {
    if (!isServicePatchingError) return null;

    return (
      <SnackAlert
        severity="error"
        message="Failed to update!"
        onClose={onCloseMessage}
      />
    );
  };

  const renderSuccessMessage = () => {
    if (!isServicePatchingSuccess) return null;

    return (
      <SnackAlert
        severity="success"
        message="The service has been updated"
        onClose={onCloseMessage}
      />
    );
  };

  const dialog = renderDialog();
  const errorMessage = renderErrorMessage();
  const successMessage = renderSuccessMessage();

  return (
    <div className="edit-group-button">
      <IconButton disabled={isServicePatching} onClick={onOpenDialog}>
        Edit Group
      </IconButton>
      {dialog}
      {errorMessage}
      {successMessage}
    </div>
  );
}
