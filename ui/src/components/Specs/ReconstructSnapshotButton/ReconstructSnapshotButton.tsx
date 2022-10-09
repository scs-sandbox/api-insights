import { ButtonHTMLAttributes, useState } from 'react';
import { Alert, Snackbar } from '@mui/material';
import IconButton from '../../IconButton/IconButton';
import BusyIcon, { BusyIconType } from '../../BusyIcon/BusyIcon';
import EnvIcon from '../SpecStateIcon/SpecStateIcon';
import { SpecData, useReconstructSnapshot } from '../../../query/spec';
import classNames from '../../../utils/className';
import { ApiError } from '../../../query/api';
import './ReconstructSnapshotButton.scss';

export type ReconstructedEvent = {
  id: string;
};

type Props = ButtonHTMLAttributes<HTMLButtonElement> & {
  serviceId?: string;
  onReconstructed?: (data: ReconstructedEvent) => void;
};

export default function ReconstructSnapshotButton(props: Props) {
  const [message, setMessage] = useState('messsage unset');
  const {
    isLoading,
    isError,
    isSuccess,
    mutate: reconstructSnapshot,
    reset: resetReconstructSnapshot,
  } = useReconstructSnapshot(props.serviceId);

  const onClick = () => {
    reconstructSnapshot(null, {
      onSuccess: (data) => {
        if (props.onReconstructed) {
          props.onReconstructed(data);
        }
        setMessage('Succeed in reconstructing snapshot');
      },
      onError: (data: ApiError) => {
        if (data.status === 400) {
          setMessage('apiclarity: no API traffic found');
        } else {
          setMessage('apiclarity not integrated for given service');
        }
      },
    });
  };

  const onCloseMessage = () => {
    resetReconstructSnapshot();
  };

  const renderMessage = () => {
    if (!isError && !isSuccess) return null;

    const severity = isSuccess ? 'success' : 'error';

    return (
      <Snackbar
        open
        anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
        autoHideDuration={3000}
        onClose={onCloseMessage}
      >
        <Alert
          onClose={onCloseMessage}
          severity={severity}
          sx={{ width: '100%' }}
        >
          {message}
        </Alert>
      </Snackbar>
    );
  };

  const renderIcon = () => (
    <i className="recons-icon">
      <EnvIcon value={SpecData.SpecState.Release} />
      <BusyIcon type={BusyIconType.ArrowCircle} busy={isLoading} />
    </i>
  );

  const className = classNames('reconstruct-snapshot-button', props.className);
  const responseMessage = renderMessage();

  return (
    <div className={className}>
      <IconButton
        icon={renderIcon()}
        disabled={props.disabled || isLoading}
        onClick={onClick}
      >
        {props.children || 'Reconstruct Snapshot'}
      </IconButton>
      {responseMessage}
    </div>
  );
}
