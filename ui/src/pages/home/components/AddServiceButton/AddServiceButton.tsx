import { ButtonHTMLAttributes, useState } from 'react';
import { useOutletContext } from 'react-router-dom';
import { ServiceData, useAddService } from '../../../../query/service';
import EditServiceDialog, {
  EditingEvent,
} from '../EditServiceDialog/EditServiceDialog';
import IconButton from '../../../../components/IconButton/IconButton';
import UploadIcon from '../../../../components/UploadIcon/UploadIcon';
import SnackAlert from '../../../../components/SnackAlert/SnackAlert';
import './AddServiceButton.scss';
import { OrganizationData } from '../../../../query/organization';

type Props = ButtonHTMLAttributes<HTMLButtonElement> & {
  authEnabled: boolean;
  organizationList: OrganizationData.Organization[];
  onServiceCreated?: (data: ServiceData.CreateServiceData) => void;
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
      onSuccess: (createServiceData: ServiceData.CreateServiceData) => {
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
