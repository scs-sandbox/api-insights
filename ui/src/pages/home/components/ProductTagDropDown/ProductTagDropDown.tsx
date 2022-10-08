import DropDownEditBox from '../../../../components/DropDown/DropDownEditBox/DropDownEditBox';
import DropDown from '../../../../components/DropDown/DropDown';

type Props = {
  value: string;
  editBoxReadOnly?: boolean;
  list: string[];
  onChange: (e: string) => void;
};

export default function ProductTagDropDown(props: Props) {
  const renderValue = () => (
    <DropDownEditBox
      value={props.value}
      readonly={props.editBoxReadOnly}
      onChange={props.onChange}
    />
  );

  const renderMenuItemLabel = (option: string) => (
    <div className="menu-item-label">{option}</div>
  );

  return (
    <DropDown
      className="product-dropdown"
      value={props.value}
      options={props.list}
      renderValue={renderValue}
      renderMenuItemLabel={renderMenuItemLabel}
      onChange={props.onChange}
    />
  );
}
