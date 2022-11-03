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

import { ButtonHTMLAttributes, useState } from 'react';
import { ServiceData, useAddService } from '../../../../query/service';
import EditServiceDialog, {
  EditingEvent,
} from '../EditServiceDialog/EditServiceDialog';
import IconButton from '../../../../components/IconButton/IconButton';
import UploadIcon from '../../../../components/UploadIcon/UploadIcon';
import SnackAlert from '../../../../components/SnackAlert/SnackAlert';
import { OrganizationData } from '../../../../query/organization';
import './AddServiceButton.scss';

type Props = ButtonHTMLAttributes<HTMLButtonElement> & {
  authEnabled: boolean;
  organizationList: OrganizationData.Organization[];
  onServiceCreated?: (data: ServiceData.Service) => void;
  tags: string[];
};

/**
 * Button that creates a new service
 * @param props object containg onServiceCreated function for service creation
 * @returns JSX.element
 */
export default function AddServiceButton(props: Props) {
  const [openDialog, setOpenDialog] = useState(false);
  const {
    isLoading: isServiceAdding,
    isError: isServiceAddingError,
    isSuccess: isServiceAddingSuccess,
    mutate: addService,
    reset: resetAddServiceResult,
  } = useAddService();

  // Opens dialog popup for new service details.
  const onOpenDialog = () => {
    setOpenDialog(true);
  };

  const onCloseDialog = () => {
    setOpenDialog(false);
  };

  const onCloseMessage = () => {
    resetAddServiceResult();
  };

  const onCreatingService = (e: EditingEvent) => {
    const data = {
      additional_info: {},
      contact: {
        email: e.email,
        name: e.contactName,
        url: e.url,
      },
      description: e.description,
      visibility: e.visibility,
      name_id: e.nameId,
      organization_id: e.organizationId,
      product_tag: e.productTag,
      title: e.title,
    };

    addService(data, {
      onSuccess: (createServiceData: ServiceData.Service) => {
        setOpenDialog(false);
        if (props.onServiceCreated) {
          props.onServiceCreated(createServiceData);
        }
      },
    });
  };

  function renderButton() {
    return (
      <IconButton
        icon={<UploadIcon busy={isServiceAdding} />}
        disabled={
          props.disabled
          || isServiceAdding
          || !props.organizationList
        }
        onClick={onOpenDialog}
      >
        {props.children || 'Add New Service'}
      </IconButton>
    );
  }

  const renderDialog = () => {
    if (!openDialog) return null;

    return (
      <EditServiceDialog
        open
        authEnabled={props.authEnabled}
        organizationList={props.organizationList}
        productTagList={props.tags}
        busy={isServiceAdding}
        handleClose={onCloseDialog}
        onEditing={onCreatingService}
      />
    );
  };

  const renderErrorMessage = () => {
    if (!isServiceAddingError) return null;

    return (
      <SnackAlert
        severity="error"
        message="Failed to create!"
        onClose={onCloseMessage}
      />
    );
  };

  const renderSuccessMessage = () => {
    if (!isServiceAddingSuccess) return null;

    return (
      <SnackAlert
        severity="success"
        message="The service has been created"
        onClose={onCloseMessage}
      />
    );
  };

  return (
    <div className="add-service-button">
      {renderButton()}
      {renderDialog()}
      {renderErrorMessage()}
      {renderSuccessMessage()}
    </div>
  );
}
