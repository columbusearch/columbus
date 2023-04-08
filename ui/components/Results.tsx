import { SearchResult } from '../types';

type Props = {
  results: SearchResult[];
};

const Results = ({ results }: Props) => {
  return (
    <ul>
      {results.map((result) => (
        <li key={result.id}>{result.title}</li>
      ))}
    </ul>
  );
};

export default Results;
