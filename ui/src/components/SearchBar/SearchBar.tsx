import './SearchBar.scss';

type Props = {
  searchKey: string;
  onSearchKeyChanged: (e) => void;
  onSearchKeyCleared: () => void;
};

export default function SearchBar(props: Props) {
  return (
    <div className="search-bar">
      <input
        type="text"
        className="search-input"
        value={props.searchKey}
        onChange={props.onSearchKeyChanged}
        placeholder="Search"
      />
      <div className="search-clear" onClick={props.onSearchKeyCleared} />
    </div>
  );
}
