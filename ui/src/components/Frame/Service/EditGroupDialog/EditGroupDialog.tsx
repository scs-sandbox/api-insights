import { useState, ChangeEvent } from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogBody from '../../../Dialog/DialogBody/DialogBody';
import FieldItem from '../../../FieldItem/FieldItem';
import DialogActions from '../../../Dialog/DialogActions/DialogActions';
import DialogTitle from '../../../Dialog/DialogTitle/DialogTitle';
import IconButton from '../../../IconButton/IconButton';
import BusyIcon from '../../../BusyIcon/BusyIcon';
import { ServiceData } from '../../../../query/service';
import './EditGroupDialog.scss';

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
  const onClose = props.busy ? null : props.handleClose;

  return (
    <Dialog open={props.open} fullWidth maxWidth="sm">
      <DialogTitle onClose={onClose}>Edit Service Details</DialogTitle>
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
