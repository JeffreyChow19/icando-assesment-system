import React, {
  createContext,
  ReactElement,
  useContext,
  useEffect,
  useState,
} from "react";
import { useQuery } from "@tanstack/react-query";
import { User } from "../interfaces/user.tsx";
import { AxiosError } from "axios";
import { removeToken } from "../utils/local-storage.ts";
import { checkAuth } from "../services/auth.ts";

interface UserContextValue {
  user: User | undefined;
  loading: boolean;
  setUser: React.Dispatch<React.SetStateAction<User | undefined>>;
  refresh: () => void;
}

export const UserContext = createContext<UserContextValue>({
  user: undefined,
  loading: true,
  // eslint-disable-next-line @typescript-eslint/no-empty-function
  setUser: () => {},
  refresh: () => {},
});

export const UserProvider = ({ children }: { children: ReactElement }) => {
  const [user, setUser] = useState<User>();
  const [loading, setLoading] = useState<boolean>(true);
  const { data, isLoading, error, refetch } = useQuery({
    queryKey: ["user"],
    queryFn: () => checkAuth(),
    retry: true,
  });

  useEffect(() => {
    if (isLoading) {
      return;
    }

    setLoading(false);
    if (error && error instanceof AxiosError && error.response?.status == 401) {
      removeToken();
      setUser(undefined);
      return;
    }

    if (!error && data) {
      setUser(data);
    }
  }, [isLoading, data, error]);

  return (
    <UserContext.Provider
      value={{ user, loading, setUser, refresh: () => refetch() }}
    >
      {children}
    </UserContext.Provider>
  );
};

export const useUser = () => useContext(UserContext);
