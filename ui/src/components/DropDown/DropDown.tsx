/*
 * Copyright 2022 Cisco Systems, Inc. and its affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

import { useState, MouseEvent, ReactNode } from 'react';
import { Menu } from '@mui/material';
import classNames from '../../utils/className';
import './DropDown.scss';

type Props = {
  id?: string;
  className?: string;
  menuItemClassName?: string;
  value?: string;
  placeholder?: string;
  options?: unknown[];
  onChange?: (value: string) => void;
  renderValue?: (value?: string) => ReactNode;
  requestOptionValue?: (option: unknown) => string;
  renderMenuItemLabel?: (option: unknown) => ReactNode;
};

/**
 * DropDown Component, examples:
 *
   function DropDown_Example1() {
    const [value, setValue] = useState('');
    const onChange = (e) => { setValue(e) };
    const options = ['hello', 'world'];
    return <DropDown value={value} options={options} onChange={onChange} />
   }

   function DropDown_Example2() {
    const [value, setValue] = useState('');
    const onChange = (e) => { setValue(e) };
    const options = [{ id: '1', text: 'hello' }, { id: '2', text: 'world' }];
    const requestOptionValue = (option: unknown) => (option as {id: string, text: string}).id;
    const renderValue = (value: string) => (
      <div className="dropdown-value">{options.find((i) => i.id === value)?.text}</div>
    );
    const renderMenuItemLabel = (option: unknown) => (
      <div>
        {(option as {id: string, text: string}).text}
      </div>
    );

    return (
      <DropDown
        value={value}
        options={options}
        requestOptionValue={requestOptionValue}
        renderValue={renderValue}
        renderMenuItemLabel={renderMenuItemLabel}
    );
  }
 */
export default function DropDown(props: Props) {
  const [anchorEl, setAnchorEl] = useState<HTMLDivElement>();

  const onOpenMenu = (e: MouseEvent<HTMLDivElement>) => {
    if (!props.options) return;

    setAnchorEl(e.currentTarget);
  };

  const onCloseMenu = () => {
    setAnchorEl(undefined);
  };

  const onChange = (e: MouseEvent<HTMLDivElement>) => {
    onCloseMenu();

    if (props.onChange) {
      props.onChange(e.currentTarget.dataset.value || '');
    }
  };

  const getOptionValue = (option: unknown) => {
    if (props.requestOptionValue) {
      return props.requestOptionValue(option);
    }

    const optionType = typeof (option);
    if (optionType === 'string') {
      return option as string;
    }

    return null;
  };

  const renderPlaceholder = () => {
    const placeholder = props.placeholder || 'Please select';

    return <div className="dropdown-placeholder">{placeholder}</div>;
  };

  const renderValue = (value?: string) => {
    if (props.renderValue) {
      return props.renderValue(value);
    }

    if (!value) return null;

    return (
      <div className="dropdown-value">{value}</div>
    );
  };

  const renderMenuItemLabel = (option: unknown) => {
    if (props.renderMenuItemLabel) {
      return props.renderMenuItemLabel(option);
    }

    const itemType = typeof (option);

    if (itemType === 'string') {
      const stringItem = option as string;

      return (
        <div className="menu-item-label">{stringItem}</div>
      );
    }

    return (
      <div className="menu-item-label">[unknown]</div>
    );
  };

  const renderMenuItems = (options?: unknown[]) => {
    if (!anchorEl) return null;

    const rect = anchorEl.getBoundingClientRect();
    const itemStyle = {
      minWidth: `${rect.width}px`,
    };

    if (!options || !options.length) {
      return (
        <div className="dropdown-menu-item" style={itemStyle} onClick={onCloseMenu}>
          <div className="menu-item-check checked" />
          <div className="menu-item-label">Please select</div>
        </div>
      );
    }

    return options
      .filter((option) => (!!option))
      .map((option, index) => {
        const optionValue = getOptionValue(option);
        const checked = props.requestOptionValue ? (
          props.requestOptionValue(option) === props.value
        ) : optionValue === props.value;
        const menuItem = renderMenuItemLabel(option);
        const key = optionValue || index;
        const dataValue = optionValue || '';

        const className = classNames('dropdown-menu-item', props.menuItemClassName);

        return (
          <div
            key={key}
            className={className}
            style={itemStyle}
            onClick={onChange}
            data-value={dataValue}
          >
            <div className={classNames('menu-item-check', checked ? 'checked' : '')} />
            {menuItem}
          </div>
        );
      });
  };

  const value = renderValue(props.value) || renderPlaceholder();
  const className = classNames('dropdown', props.className);
  const menuItems = renderMenuItems(props.options);
  const dropMenu = anchorEl ? (
    <Menu open anchorEl={anchorEl} onClose={onCloseMenu}>
      {menuItems}
    </Menu>
  ) : null;

  return (
    <div id={props.id} className={className}>
      <div className="dropdown-value-box" onClick={onOpenMenu}>
        {value}
        <div className="drop-icon" />
      </div>
      {dropMenu}
    </div>
  );
}
