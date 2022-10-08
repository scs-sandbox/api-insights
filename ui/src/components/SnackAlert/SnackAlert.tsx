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
