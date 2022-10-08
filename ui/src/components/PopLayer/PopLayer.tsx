import { useEffect, useState, HTMLAttributes } from 'react';
import { createPortal } from 'react-dom';
import './PopLayer.scss';

type Props = HTMLAttributes<HTMLElement>;

export default function PopLayer(props: Props) {
  const [popLayer, setPopLayer] = useState(
    document.querySelector('#pop-layer'),
  );

  useEffect(() => {
    if (popLayer) return;

    const domPopLayer = document.createElement('div');
    domPopLayer.id = 'pop-layer';
    document.body.appendChild(domPopLayer);

    setPopLayer(domPopLayer);
  }, []);

  return popLayer ? createPortal(props.children, popLayer) : null;
}
