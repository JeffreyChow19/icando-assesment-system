import { Link, useLocation, useNavigate } from "react-router-dom";
import { removeToken } from "../utils/local-storage.ts";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@repo/ui/components/ui/dropdown-menu";
import {
  BookOpenCheck,
  HomeIcon,
  LogOut,
  MenuIcon,
  UsersRound,
} from "lucide-react";
import {
  Sheet,
  SheetContent,
  SheetTrigger,
} from "@repo/ui/components/ui/sheet";
import { Button } from "@repo/ui/components/ui/button";
import { cn } from "@repo/ui/lib/utils";
import React from "react";
import { Avatar, AvatarFallback } from "@ui/components/ui/avatar.tsx";
import { useUser } from "../context/user-context.tsx";

interface NavItemLink {
  icon: React.ReactElement;
  title: string;
  link: string;
}

const iconClassName = "w-4 h-4";

const navItems: NavItemLink[] = [
  {
    icon: <HomeIcon className={iconClassName} />,
    title: "Dashboard",
    link: "/",
  },
  {
    icon: <BookOpenCheck className={iconClassName} />,
    title: "Quizzes",
    link: "/quiz",
  },
  {
    icon: <UsersRound className={iconClassName} />,
    title: "Students",
    link: "/student",
  },
];

const UserDropdown = () => {
  const navigate = useNavigate();
  // const { user, setUser, refresh } = useUser();
  const logout = () => {
    removeToken();
    // setUser(undefined);
    // refresh();
    navigate("/login");
  };

  const { user } = useUser();

  return (
    <DropdownMenu>
      <DropdownMenuTrigger className="items-center whitespace-nowrap rounded-md text-sm text-muted-foreground font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-primary-foreground/20 hover:text-primary-foreground p-1">
        <Avatar className="size-8">
          <AvatarFallback className="bg-secondary text-secondary-foreground font-bold">
            {user?.firstName[0].toUpperCase()}
            {user?.lastName ? user.lastName[0].toUpperCase() : ""}
          </AvatarFallback>
        </Avatar>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56 mr-4">
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
  const pathname = "/" + location.pathname.split("/")[1];

  return (
    <div className="px-2 lg:p-6 flex flex-col gap-3">
      {navItems.map((item) => {
        return (
          <Link to={item.link} key={item.link}>
            <div
              className={cn(
                pathname === item.link
                  ? "bg-primary/10 font-medium text-primary"
                  : "hover:bg-muted hover:rounded-md hover:text-primary",
                "flex flex-row gap-2 items-center pr-20 py-2 pl-3 rounded-md",
              )}
            >
              {item.icon}
              <p className="whitespace-nowrap text-md">{item.title}</p>
            </div>
          </Link>
        );
      })}
    </div>
  );
};

export const SideBar = () => {
  return (
    <div className="sticky top-0 left-0 min-h-screen max-h-screen bg-white hidden lg:block shadow-lg p-2">
      <img src={"/logo.png"} alt={"logo"} className="w-48 m-auto pt-4" />
      <NavMenus />
    </div>
  );
};

export const Navigation = () => {
  const { user } = useUser();
  return (
    <header className="flex flex-row sticky top-0 bg-primary backdrop-blur justify-between lg:justify-end px-2 lg:px-6 py-2 z-20 items-center">
      <div className="lg:hidden">
        <Sheet>
          <SheetTrigger asChild>
            <Button
              variant="ghost"
              className="hover:bg-primary-foreground/20 hover:text-primary-foreground"
            >
              <MenuIcon className="h-4 w-4 text-background" />
            </Button>
          </SheetTrigger>
          <SheetContent side="left" className="flex flex-col gap-4 pt-10">
            <NavMenus />
          </SheetContent>
        </Sheet>
      </div>
      <div className="flex items-center gap-2">
        <p className="text-primary-foreground">
          {user?.firstName + " " + user?.lastName || "-"}
        </p>
        <UserDropdown />
      </div>
    </header>
  );
};
