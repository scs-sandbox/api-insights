import { ChangeEvent, MouseEvent } from 'react';
import './DropDownEditBox.scss';

type Props = {
  value?: string;
  readonly?: boolean;
  onChange?: (value: string) => void;
};

export default function DropDownEditBox(props: Props) {
  const onChangeText = (e: ChangeEvent<HTMLInputElement>) => {
    if (props.onChange) {
      props.onChange(e.target.value);
    }
  };

  const onClick = (e: MouseEvent<HTMLInputElement>) => {
    if (props.readonly) return;

    e.stopPropagation();
  };

  return (
    <div className="dropdown-value">
      <input
        className="dropdown-value-input"
        readOnly={props.readonly}
        value={props.value}
        onChange={onChangeText}
        onClick={onClick}
      />
    </div>
  );
}
