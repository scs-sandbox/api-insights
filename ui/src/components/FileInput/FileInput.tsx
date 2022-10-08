import { InputHTMLAttributes } from 'react';
import './FileInput.scss';

type Props = InputHTMLAttributes<HTMLInputElement> & {
  action?: string;
};

export default function FileInput(props: Props) {
  const buttonClassName = `file-input ${props.className || ''}`.trim();

  return (
    <div id={props.id} className={buttonClassName}>
      <div className="input-value">{props.value}</div>
      <div className="input-action blue-block">{props.action || 'Upload'}</div>
      <label className="file-input-box">
        <input
          className="sys-input"
          type="file"
          value=""
          accept={props.accept}
          disabled={props.disabled}
          onChange={props.onChange}
        />
      </label>
    </div>
  );
}
