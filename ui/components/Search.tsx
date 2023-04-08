import { useState } from 'react';

const Search = () => {
  const [searchTerm, setSearchTerm] = useState('');

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(event.target.value);
  };

  const handleFormSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    // Add code here to handle the form submission
  };

  return (
    <form onSubmit={handleFormSubmit}>
      <input
        type="text"
        placeholder="Search"
        value={searchTerm}
        onChange={handleInputChange}
        className="px-4 py-2 text-gray-700 border rounded-full focus:outline-none focus:border-purple-500"
      />
      <button type="submit" className="ml-4 px-4 py-2 text-white bg-purple-500 rounded-full hover:bg-purple-600 focus:outline-none focus:bg-purple-600">
        Search
      </button>
    </form>
  );
};

export default Search;
