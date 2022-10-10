import { SpecData, useFetchSpecList } from '../../../query/spec';
import DropDown from '../../DropDown/DropDown';
import './SepcDropDown.scss';

type Props = {
  serviceId: string;
  selectedSpec?: SpecData.Spec;
  onChange: (spec: SpecData.Spec) => void;
};

export default function SpecDropDown(props: Props) {
  const { data, isLoading } = useFetchSpecList(props.serviceId);
  const specList = data as SpecData.Spec[];

  const onChange = (value: string) => {
    const spec = data.find((i) => i.id === value);
    if (props.onChange) {
      props.onChange(spec);
    }
  };

  const requestOptionValue = (option: SpecData.Spec) => option.id;

  const renderValue = (value: string) => {
    const spec = (specList || []).find((i) => i.id === value);

    if (!spec) return null;

    return (
      <div className="dropdown-value">
        <span className="version-label">{spec.version}</span>
        <span className="revision-label">{spec.revision}</span>
      </div>
    );
  };

  const renderMenuItemLabel = (option: SpecData.Spec) => (
    <div className="menu-item-label">
      <span className="version-label">{option.version}</span>
      <span className="revision-label">{option.revision}</span>
    </div>
  );

  const placeholder = isLoading ? 'Loading...' : 'Please select';
  const value = props.selectedSpec ? props.selectedSpec.id : '';

  return (
    <DropDown
      className="spec-dropdown"
      menuItemClassName="spec-menu-item"
      placeholder={placeholder}
      value={value}
      options={data}
      requestOptionValue={requestOptionValue}
      renderValue={renderValue}
      renderMenuItemLabel={renderMenuItemLabel}
      onChange={onChange}
    />
  );
}
