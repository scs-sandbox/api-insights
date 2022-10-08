import { useEffect, ReactElement } from 'react';

type Props = {
  ms: number; // timeout
  onClose?: () => void;
  children: ReactElement;
};

export default function Timer(props: Props) {
  useEffect(() => {
    if (!props.ms || !props.onClose) return null;

    const timer = setTimeout(() => {
      props.onClose();
    }, props.ms);

    return () => {
      clearTimeout(timer);
    };
  }, []);

  return props.children;
}
