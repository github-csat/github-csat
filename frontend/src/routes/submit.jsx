import { useSearchParams } from "react-router-dom";
import { useEffect } from "react";

export default function Submit() {
  const [searchParams, setSearchParams] = useSearchParams();
  const [error, setError] = useState();
  const [isLoggedIn, setIsLoggedIn] = useState();

  const issueUrl = searchParams.get("issueUrl");
  const feedback = searchParams.get("feedback");
  // ping session endpoint to check? Can we cache the auth state in local storage?
  // how do we clear that state on an auth failure to restart the login flow?

  // We can't store the actual session ID in localstorage for sec reasons, need to
  // use an httpOnly cookie
  //
  // JWT contents: GHUsername, signed by server
  const loggedIn = true;

  useEffect(() => {});

  useEffect(() => {
    if (loggedIn) {
      // submit the feedback, show success
      // or error in render JSX
    } else {
      // redirect to API login endpoint
      // with state parameters for submit on return
      // state = {issueUrl, feedback, path: "/submit"}
    }
  });

  return (
    <div>
      Your feedback was submitted!
      <p>Issue: {issueUrl}</p>
      <p>Feedback: {feedback}</p>
    </div>
  );
}
