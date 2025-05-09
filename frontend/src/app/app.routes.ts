import { Routes } from "@angular/router";
import { HomeComponent } from "./home/home.component";

export const routes: Routes = [
  { path: "", component: HomeComponent },
  {
    path: "tournaments",
    loadComponent: () =>
      import("./tournaments/tournaments.component").then(
        (mod) => mod.TournamentsComponent
      ),
  },
  {
    path: "tournaments/:id",
    loadComponent: () =>
      import("./tournaments/tournament/tournament.component").then(
        (mod) => mod.TournamentComponent
      ),
  },
  {
    path: "new-tournament",
    loadComponent: () =>
      import(
        "./tournaments/new-tournament-form/new-tournament-form.component"
      ).then((mod) => mod.NewTournamentFormComponent),
  },
];
