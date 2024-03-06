<!-- resources/views/siswas/index.blade.php -->

<h1>Siswas</h1>
<a href="{{ route('siswas.create') }}">Create Siswa</a>

<ul>
    @foreach($siswas as $siswa)
        <li>
            <a href="{{ route('siswas.show', $siswa->id) }}">{{ $siswa->nama }}</a>
            <a href="{{ route('siswas.edit', $siswa->id) }}">Edit</a>
            <form action="{{ route('siswas.destroy', $siswa->id) }}" method="post">
                @csrf
                @method('DELETE')
                <button type="submit">Delete</button>
            </form>
        </li>
    @endforeach
</ul>
