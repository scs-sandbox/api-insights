import { ReactNode } from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogTitle from '@mui/material/DialogTitle';
import './HelpDialog.scss';

type Props = {
  open?: boolean;
  title?: ReactNode;
  message?: ReactNode;
  handleClose?: () => void;
};

export default function HelpDialog(props: Props) {
  return (
    <Dialog
      open={props.open}
      onClose={props.handleClose}
      fullWidth
      maxWidth="sm"
    >
      <DialogTitle>{props.title}</DialogTitle>
      <div className="help-dialog-content">{props.message}</div>
      <DialogActions>
        <Button onClick={props.handleClose}>OK</Button>
        {/* <Button disabled={disabledOK} onClick={onOK}>Create</Button> */}
      </DialogActions>
    </Dialog>
  );
}
