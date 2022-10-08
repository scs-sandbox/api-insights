import classNames from '../../utils/className';
import './Toggle.scss';

type Props = {
  checked?: boolean;
  label?: string;
  onToggle?: (checked: boolean) => void;
};

export default function Toggle(props: Props) {
  const label = props.label ? (
    <div className="toggle-label">{props.label}</div>
  ) : null;
  const className = classNames('toggle', props.checked ? 'checked' : '');

  const onToggle = () => {
    if (props.onToggle) {
      props.onToggle(!props.checked);
    }
  };

  return (
    <div className={className} onClick={onToggle}>
      <div className="toggle-switch">
        <div className="switch-button" />
      </div>
      {label}
    </div>
  );
}
