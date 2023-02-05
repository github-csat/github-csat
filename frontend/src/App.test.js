import { render, screen } from '@testing-library/react';
import Satisfactions from './Satisfactions';

test('renders learn react link', () => {
  render(<Satisfactions />);
  const linkElement = screen.getByText(/learn react/i);
  expect(linkElement).toBeInTheDocument();
});
