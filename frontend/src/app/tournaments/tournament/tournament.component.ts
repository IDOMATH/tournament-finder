import { Component, inject, signal } from "@angular/core";
import { TournamentService } from "../../services/tournament-service";
import { Tournament } from "../../models/tournament.model";
import { DatePipe } from "@angular/common";
import { AgeDivisionComponent } from "../../age-division/age-division.component";

@Component({
  selector: "app-tournament",
  imports: [DatePipe, AgeDivisionComponent],
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
