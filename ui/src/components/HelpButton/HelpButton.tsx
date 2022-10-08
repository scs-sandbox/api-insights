import { useState, useEffect, ReactNode } from 'react';
import IconButton from '../IconButton/IconButton';
import TalkIcon from './images/talk.png';
import './HelpButton.scss';

type Props = {
  show?: boolean;
  title: ReactNode;
  message: ReactNode;
};

export default function HelpButton(props: Props) {
  const [openDialog, setOpenDialog] = useState(props.show);

  useEffect(() => {
    setOpenDialog(props.show);
  }, [props.show]);

  function renderButton() {
    return (
      <IconButton
        // icon={<HelpIcon/>}
        onClick={() => setOpenDialog(!openDialog)}
      >
        {/* {props.children || 'Help'} */}
        ?
      </IconButton>
    );
  }

  function renderBar() {
    if (!openDialog) return null;

    return (
      <div className="help-bar">
        <img className="talk-icon" alt="help" src={TalkIcon} />
        <div className="title">{props.title}</div>
        <div className="description">{props.message}</div>
      </div>
    );
  }

  const button = renderButton();
  const bar = renderBar();

  return (
    <div className={`help-container ${openDialog ? 'show' : ''}`}>
      <div className="help-button">
        {button}
      </div>
      {bar}
    </div>
  );
}
