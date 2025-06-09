import ChannelChat from "../../../../components/ChannelChat";

export default function Page({ page }) {
    const { serverId, channelId } = params;
    return <ChannelChat serverId={serverId} channelId={channelId}/>;
}
