
import { useSearchParams } from 'react-router-dom';

export default function LoggedIn() {
    const [searchParams, setSearchParams] = useSearchParams();
    return <div>You're logged in as {searchParams.get("name")} ({searchParams.get("handle")})</div>
}