import { vi, describe, test, expect, afterEach } from "vitest";
import Login from "./login-config";
import { render, screen } from "@testing-library/react";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import * as reactRouterDom from "react-router-dom";

describe("Post-Login test", () => {
  // need to render in a router so that we can
  // access querystring params
  const router = createBrowserRouter([
    {
      path: "/",
      element: <Login />,
    },
  ]);

  afterEach(() => {
    vi.resetAllMocks();
  });

  test("always shows greeting", () => {
    render(<RouterProvider router={router} />);
    expect(screen.getByText(/What's Shakin/i)).toBeDefined();
  });
  test("gets params from querystring", () => {
    // this is progably dumb but it works for now
    vi.spyOn(reactRouterDom, "useSearchParams").mockReturnValue([
      {
        get: (key) => {
          if (key === "name") {
            return "Dex";
          } else if (key === "handle") {
            return "dexhorthy";
          }
        },
      } as any,
      null,
    ]);

    render(<RouterProvider router={router} />);
    expect(screen.getByText(/What's Shakin' Dex \(dexhorthy\)/i)).toBeDefined();
  });
});
