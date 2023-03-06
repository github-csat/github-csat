
import { useSearchParams } from 'react-router-dom';

export default function LoggedIn() {
    const [searchParams, setSearchParams] = useSearchParams();
    let name = searchParams.get("name");
    let handle = searchParams.get("handle");
    return <div>What's shakin' {name} ({handle})</div>
}