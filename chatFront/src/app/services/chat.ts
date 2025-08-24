import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ChatService {
  private ws!: WebSocket;
  private messagesSubject = new Subject<string>();
  public messages$ = this.messagesSubject.asObservable();

  connect(username: string) {
    this.ws = new WebSocket(`${window.location.origin}/ws`);

    this.ws.onopen = () => {
      console.log('Connected to WebSocket');
      this.ws.send(username); // first message = username
    };

    this.ws.onmessage = (event) => {
      this.messagesSubject.next(event.data);
    };

    this.ws.onclose = () => console.log('WebSocket closed');
    this.ws.onerror = (err) => console.error('WebSocket error', err);
  }

  sendMessage(msg: string) {
    if (this.ws && this.ws.readyState === WebSocket.OPEN) {
      this.ws.send(msg);
    }
  }

  async getOnlineCount(): Promise<number> {
    try {
      const res = await fetch(`${window.location.origin}/online`);
      if (!res.ok) throw new Error('network');
      const j = await res.json();
      return j?.online ?? 0;
    } catch (e) {
      console.error('Failed to fetch online count', e);
      return 0;
    }
  }
}
