import { Component } from "@angular/core";
import { RouterOutlet } from "@angular/router";
import { NewTournamentFormComponent } from "./tournaments/new-tournament-form/new-tournament-form.component";
import { HeaderComponent } from "./header/header.component";

@Component({
  selector: "app-root",
  imports: [RouterOutlet, NewTournamentFormComponent, HeaderComponent],
  templateUrl: "./app.component.html",
  styleUrl: "./app.component.css",
})
export class AppComponent {
  title = "tournament-finder";
}
