import { OrganizationData } from '../../../../query/organization';
import DropDownEditBox from '../../../../components/DropDown/DropDownEditBox/DropDownEditBox';
import DropDown from '../../../../components/DropDown/DropDown';

type Props = {
  value: string;
  editBoxReadOnly?: boolean;
  list: OrganizationData.Organization[];
  onChange: (e: string) => void;
};

export default function OrganizationDropDown(props: Props) {
  const requestOptionValue = (option: OrganizationData.Organization) => option.name_id;

  const renderValue = () => {
    const foundOrganizationItem = (props.list || []).find(
      (i) => requestOptionValue(i) === props.value,
    );

    const value = foundOrganizationItem ? foundOrganizationItem.title : props.value;

    return (
      <DropDownEditBox value={value} readonly={props.editBoxReadOnly} onChange={props.onChange} />
    );
  };

  const renderMenuItemLabel = (option: OrganizationData.Organization) => (
    <div className="menu-item-label">{option.title}</div>
  );

  return (
    <DropDown
      className="organization-dropdown"
      value={props.value}
      options={props.list}
      renderValue={renderValue}
      requestOptionValue={requestOptionValue}
      renderMenuItemLabel={renderMenuItemLabel}
      onChange={props.onChange}
    />
  );
}
