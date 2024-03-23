export interface User {
  id: string;
  email: string;
  firstName: string;
  lastName: string;
}

export interface Teacher extends User {
  institutionId: string;
}
