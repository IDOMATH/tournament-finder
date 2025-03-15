import { Injectable } from "@angular/core";
import { Tournament } from "./tournament.model";

@Injectable({ providedIn: "root" })
export class TournamentService {
  getTournamentById(id: number) {}
  getTournaments() {}
  postTournament(tournament: Tournament) {
    JSON.stringify(tournament);
    // call to post tournament
  }
  putTournament(tournament: Tournament) {
    JSON.stringify(tournament);
  }
  deleteTournament(id: number) {}
}
