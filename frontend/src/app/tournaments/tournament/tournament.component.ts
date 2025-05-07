import { Component, inject, signal } from "@angular/core";
import { TournamentService } from "../tournament-service";
import { Tournament } from "../tournament.model";

@Component({
  selector: "app-tournament",
  imports: [],
  templateUrl: "./tournament.component.html",
  styleUrl: "./tournament.component.css",
})
export class TournamentComponent {
  private tournamentService = inject(TournamentService);
  tournament = signal<Tournament | null>(null);

  ngOnInit() {
    this.tournament.set(this.tournamentService.getTournamentById(17));
  }
}
