import { Link, useLocation, useNavigate } from "react-router-dom";
import { removeToken } from "../utils/local-storage.ts";
import { useUser } from "../context/user-context.tsx";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@repo/ui/components/ui/dropdown-menu";
import { HomeIcon, LogOut, MenuIcon, UserRound } from "lucide-react";
import {
  Sheet,
  SheetContent,
  SheetTrigger,
} from "@repo/ui/components/ui/sheet";
import { Button } from "@repo/ui/components/ui/button";
import { cn } from "@repo/ui/lib/utils";
import React from "react";

interface NavItemLink {
  icon: React.ReactElement;
  title: string;
  link: string;
}

const iconClassName = "w-4 h-4";

const navItems: NavItemLink[] = [
  {
    icon: <HomeIcon className={iconClassName} />,
    title: "Home",
    link: "/",
  },
];

const UserDropdown = () => {
  const navigate = useNavigate();
  const { studentQuiz: user, setStudentQuiz: setUser } = useUser();
  const logout = () => {
    removeToken();
    setUser(undefined);
    navigate("/login");
  };

  return (
    <DropdownMenu>
      <DropdownMenuTrigger className="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-primary-foreground/20 hover:text-primary-foreground h-10 px-4 py-2">
        <p className={"text-white mr-2"}>Hello, {user?.name}</p>
        <UserRound className="h-4 w-4 text-white" />
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56">
        <DropdownMenuItem onClick={logout}>
          <LogOut className="mr-2 h-4 w-4" />
          <span>Logout</span>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

const NavMenus = () => {
  const location = useLocation();
  const pathname = location.pathname;
  // const { user } = useUser();

  return (
    <div className="p-2 lg:p-6 flex flex-col gap-2">
      {navItems.map((item) => {
        return (
          <Link to={item.link} key={item.link}>
            <div
              className={cn(
                pathname === item.link
                  ? "bg-primary/5"
                  : "hover:bg-transparent hover:underline",
                "flex flex-row gap-2 text-sm items-center pr-20 py-1 pl-2 rounded-md",
              )}
            >
              {item.icon}
              <p className="whitespace-nowrap">{item.title}</p>
            </div>
          </Link>
        );
      })}
    </div>
  );
};

export const SideBar = () => {
  return (
    <div className="sticky top-0 left-0 min-h-full bg-primary-foreground hidden lg:block">
      <NavMenus />
    </div>
  );
};

export const Navigation = () => {
  return (
    <header className="flex flex-row sticky top-0 bg-foreground/95 backdrop-blur justify-between lg:justify-end px-2 lg:px-6 py-2 z-20 items-center">
      <div className="lg:hidden">
        <Sheet>
          <SheetTrigger asChild>
            <Button
              variant="ghost"
              className="hover:bg-primary-foreground/20 hover:text-primary-foreground"
            >
              <MenuIcon className="h-4 w-4 text-white" />
            </Button>
          </SheetTrigger>
          <SheetContent side="left" className="flex flex-col gap-4">
            <NavMenus />
          </SheetContent>
        </Sheet>
      </div>
      <UserDropdown />
    </header>
  );
};
