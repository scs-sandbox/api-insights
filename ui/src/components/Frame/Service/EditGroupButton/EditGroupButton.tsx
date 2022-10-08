import { useState, ButtonHTMLAttributes } from 'react';
import { ServiceData, usePatchService } from '../../../../query/service';
import IconButton from '../../../IconButton/IconButton';
import SnackAlert from '../../../SnackAlert/SnackAlert';
import EditGroupDialog, {
  EditingEvent,
} from '../EditGroupDialog/EditGroupDialog';
import './EditGroupButton.scss';

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
