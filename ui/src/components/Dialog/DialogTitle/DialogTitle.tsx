import { ReactNode } from 'react';
import IconButton from '@mui/material/IconButton';
import CloseIcon from '@mui/icons-material/Close';
import SettingsIcon from '@mui/icons-material/Settings';
import { DialogTitle as MuiDialogTitle } from '@mui/material';
import './DialogTitle.scss';

type Props = {
  icon?: ReactNode;
  children?: ReactNode;
  onClose?: () => void;
};

export default function DialogTitle(props: Props) {
  const {
    icon, children, onClose, ...other
  } = props;

  const sx = {
    position: 'absolute',
    right: 8,
    top: 8,
    color: (theme) => theme.palette.grey[500],
  };

  return (
    <MuiDialogTitle className="dialog-title" sx={{ m: 0, p: 2 }} {...other}>
      {icon || <SettingsIcon className="title-icon" />}
      {children}
      <IconButton aria-label="close" onClick={onClose} sx={sx}>
        <CloseIcon />
      </IconButton>
    </MuiDialogTitle>
  );
}
